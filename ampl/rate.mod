param PI := 3.141592653589793238462643383279502884197;
param N_REAL_ESTATES;
param N_BUSINESSES;

set RealEstates = {1..N_REAL_ESTATES};
set RealEstatesOmitted default {};
set Businesses = {1..N_BUSINESSES};

param REAL_ESTATE_LNG {RealEstates};
param REAL_ESTATE_LAT {RealEstates};
param CRIME_RATIO {RealEstates};

param BUSINESS_LNG {Businesses};
param BUSINESS_LAT {Businesses};

param DIST {r in RealEstates, b in Businesses}
  = acos(
     sin(REAL_ESTATE_LAT[r]/180*PI)*sin(BUSINESS_LAT[b]/180*PI)
     +cos(REAL_ESTATE_LAT[r]/180*PI)*cos(BUSINESS_LAT[b]/180*PI)*cos((BUSINESS_LNG[b]-REAL_ESTATE_LNG[r])/180*PI)
   ) * 6371;

param DIST_MAX {r in RealEstates} = max {b in Businesses} DIST[r, b];

var chosen {r in RealEstates} binary;
s.t. chosen_only_one:
  sum{r in RealEstates} chosen[r] = 1;
s.t. chosen_not_omitted {r in RealEstates : r in RealEstatesOmitted}:
  chosen[r] = 0;

set Targets = {"dist", "crime_ratio"};
var dist_score >= 0;
var crime_ratio_score >= 0;
s.t. dist_score_ge_chosen {r in RealEstates}:
  dist_score >= DIST_MAX[r] * chosen[r];
s.t. crime_ratio_score_ge_chosen {r in RealEstates}:
  crime_ratio_score >= CRIME_RATIO[r] * chosen[r];
  
param EPSILON default 1e-10;
param BETA default 1e-5;
param ASPIRATION {Targets} default 0;
param LAMBDA {Targets} default 1;

var target_values {Targets};
var target_min;

s.t. target_dist:
  target_values["dist"] <= -LAMBDA["dist"] * (dist_score - ASPIRATION["dist"]);

s.t. target_dist_beta:
  target_values["dist"] <= -BETA * LAMBDA["dist"] * (dist_score - ASPIRATION["dist"]);

s.t. target_crime_ratio:
  target_values["crime_ratio"] <= -LAMBDA["crime_ratio"] * (crime_ratio_score - ASPIRATION["crime_ratio"]);

s.t. target_crime_ratio_beta:
  target_values["crime_ratio"] <= -BETA * LAMBDA["crime_ratio"] * (crime_ratio_score - ASPIRATION["crime_ratio"]);


s.t. target_min_less_than_others {t in Targets}:
  target_min <= target_values[t];

maximize reference_point_criterium:
  target_min + BETA * sum {t in Targets} target_values[t];

