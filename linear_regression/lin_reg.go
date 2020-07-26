package main

//Checkout the GoLearn Repo: https://github.com/sjwhitworth/golearn

import (
	"fmt"

	"github.com/sjwhitworth/golearn/linear_models"
	"github.com/sjwhitworth/golearn/base"
)

func main() {

	//Parese CSVinto GoLearn Instances
	rawData, err := base.ParseCSVToInstances("datasets/exams.csv", true)
	if err != nil {
		panic(err)
	}

	//Split data: 70% -> Training, 30% -> Test
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.70)

	//Instantiate new linear regression model struct
	linReg := linear_models.NewLinearRegression()

	//Fit model
	linReg.Fit(trainData)

	//Generate predictions
	predictions, err := linReg.Predict(testData)

	//Compare results
	//TODO: Generate MSE (Mean squared error) to simplify performance assessment.
	fmt.Print(testData)
	fmt.Print(predictions)

}
