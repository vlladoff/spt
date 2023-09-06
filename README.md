# SPT
Installation
To use SPT, follow these steps:

1. Clone the repository to your local machine:

   ```shell
   git clone https://github.com/vlladoff/spt.git

2. Run (example):

   ```shell
   go run cmd/spt/main.go -source='/test_data.csv'


## Usage
SPT accepts the following command-line options:

````
-model (default: "ext"): Specify the prediction model to use. You can choose between "ext" (linear extrapolation) and "reg" (linear regression).

-source: Specify the path to the source file containing the data you want to analyze.

-aggregate (default: "country"): Choose the aggregation type for the analysis. You can aggregate the data by "country" or "campaign."

-sort (default: "name"): Specify the sorting type for the results. You can sort by "name" or "value."