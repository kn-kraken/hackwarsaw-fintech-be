package api

import (
	"fmt"
	"io"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/kn-kraken/hackwarsaw-fintech/lib/db"
)

func Rate(businesses []db.Business, realEstates []db.RealEstate) ([]float32, error) {
	dir, err := os.MkdirTemp("", "kraken-fintech-ampl")
	if err != nil {
		return nil, fmt.Errorf("making temp dir: %w", err)
	}
	defer os.RemoveAll(dir)

	cp := exec.Command("cp", "-r", "./ampl", dir)
	if err := cp.Run(); err != nil {
		return nil, fmt.Errorf("copying ampl files: %w", err)
	}

	dat, err := os.Create(dir + "/ampl/rate.dat")
	if err != nil {
		return nil, fmt.Errorf("creating rate.dat: %w", err)
	}
	defer dat.Close()

	err = writeAmplData(dat, businesses, realEstates)
	if err != nil {
		return nil, fmt.Errorf("writing rate.dat: %w", err)
	}

	ampl := exec.Command("ampl", "./ampl/rate.run")
	ampl.Dir = dir
	out, err := ampl.Output()
	if err != nil {
		slog.Info("ampl output", "output", out)
		return nil, fmt.Errorf("running ampl command: %w", err)
	}
	scores := readAmplOut(out)

	return scores, nil
}

func writeAmplData(writer io.Writer, businesses []db.Business, realEstates []db.RealEstate) error {
	fmt.Fprintf(writer, "param N_BUSINESSES = %v;\n", len(businesses))
	fmt.Fprintf(writer, "param N_REAL_ESTATES = %v;\n", len(realEstates))

	fmt.Fprintf(writer, "param : REAL_ESTATE_LNG, REAL_ESTATE_LAT, CRIME_RATIO = \n")
	for i, realEstate := range realEstates {
		fmt.Fprintf(
			writer, "%d %v, %v, %v\n",
			i+1, realEstate.Location.Longitude, realEstate.Location.Latitude, rand.Float32(),
		)
	}
	fmt.Fprintf(writer, ";\n")

	fmt.Fprintf(writer, "param : BUSINESS_LNG, BUSINESS_LAT = \n")
	for i, business := range businesses {
		fmt.Fprintf(
			writer, "%d %v, %v\n",
			i+1, business.Location.Longitude, business.Location.Latitude,
		)
	}
	fmt.Fprintf(writer, ";\n")

	return nil
}

func readAmplOut(out []byte) []float32 {
	min := math.Inf(1)
	max := math.Inf(-1)
	var scores []float32

	lines := strings.Split(string(out), "\n")[1:]
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == ";" {
			break
		}
		fields := strings.Fields(line)
		score, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			panic(err)
		}

		if score > max {
			max = score
		}
		if score < min {
			min = score
		}

		scores = append(scores, float32(score))
	}

	for i, score := range scores {
		scores[i] = float32((float64(score) - min) / (max - min))
	}

	return scores
}

func copy(src string, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
