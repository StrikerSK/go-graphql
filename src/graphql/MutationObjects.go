package graphql

import (
	"github.com/graphql-go/graphql"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createMax": &graphql.Field{
			Type:        exerciseMax,
			Description: "Create exercise max",
			Args: graphql.FieldConfigArgument{
				"Max": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"Date": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"ExerciseID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				max, _ := params.Args["Max"].(float64)
				date, _ := params.Args["Date"].(string)
				exerciseID, _ := params.Args["ExerciseID"].(int)

				newMax := entity.ExerciseMaxPostRequest{
					Max:          float32(max),
					ExerciseID:   exerciseID,
					TimeExecuted: date,
				}

				service.CreateExercisePersonalBest(newMax)
				return service.ReadExercisePersonalBest(exerciseID), nil
			},
		},
		"createTemplate": &graphql.Field{
			Type:        exerciseTemplate,
			Description: "Create exercise template",
			Args: graphql.FieldConfigArgument{
				"Plan": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.NewList(inputPlan)),
				},
				"Note": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"ExerciseID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				plan, _ := params.Args["Plan"].([]interface{})
				note, _ := params.Args["Note"].(string)
				exerciseID, _ := params.Args["ExerciseID"].(int)

				templatePlan := entity.ExerciseTemplatePlan{
					Plan:      entity.TemplatePlan{},
					WorkoutId: exerciseID,
					Note:      note,
				}

				enrichTemplate(&templatePlan, plan)
				service.CreateTemplate(templatePlan)
				return service.ReadTemplate(exerciseID), nil
			},
		},
		"createHistory": &graphql.Field{
			Type:        exerciseHistory,
			Description: "Create exercise history data",
			Args: graphql.FieldConfigArgument{
				"Date": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"Sets": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.NewList(inputPlan)),
				},
				"Notes": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"ExerciseID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				sets, _ := params.Args["Sets"].([]interface{})
				note, _ := params.Args["Notes"].(string)
				date, _ := params.Args["Date"].(string)
				exerciseID, _ := params.Args["ExerciseID"].(int)

				newHistoryEntry := entity.PostExerciseHistory{
					Date:      date,
					Plan:      processSetPlan(sets),
					WorkoutId: exerciseID,
					Notes:     note,
				}

				service.CreateExerciseHistory(newHistoryEntry)
				return newHistoryEntry, nil
			},
		},
	},
})

func enrichTemplate(input *entity.ExerciseTemplatePlan, outputSet []interface{}) {
	var sets []entity.ExerciseTemplateSet

	for _, firstIt := range outputSet {
		var templateSet entity.ExerciseTemplateSet
		for _, tempItem := range firstIt.([]interface{}) {
			templateVolume := entity.ExerciseTemplateVolume{
				ExerciseRepsAndWeight: extractRep(tempItem),
			}
			templateSet.Repetitions = append(templateSet.Repetitions, templateVolume)
		}
		sets = append(sets, templateSet)
	}

	if sets != nil {
		input.Plan = sets
	}

	return
}

func processSetPlan(sets []interface{}) (output [][]entity.ExerciseRepsAndWeight) {
	for _, set := range sets {
		var processedSet []entity.ExerciseRepsAndWeight
		for _, item := range set.([]interface{}) {
			processedSet = append(processedSet, extractRep(item))
		}
		output = append(output, processedSet)
	}
	return
}

func extractRep(item interface{}) entity.ExerciseRepsAndWeight {
	repsAndWeigh := item.(map[string]interface{})
	reps := repsAndWeigh["Reps"].(int)
	weight := repsAndWeigh["Weight"].(float64)
	output := entity.ExerciseRepsAndWeight{
		Reps:   reps,
		Weight: float32(weight),
	}
	return output
}

//For creating input object
var inputTodo = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Todo",
		Fields: graphql.InputObjectConfigFieldMap{
			"Name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"Done": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	},
)
