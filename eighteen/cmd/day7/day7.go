package main

import (
	"eighteen/internal/parser"
	"fmt"
	"sort"
	"strings"
)

const FinishedStep = "."
const InProgressStep = "*"
const SecondsForEachStep = 60

func main() {
	strnarr := parser.ParseStrings("day7.txt")
	// EFHLMTKQBWAPGIVXSZJRDUYONC
	//result := part1(strnarr)
	//if result != "EFHLMTKQBWAPGIVXSZJRDUYONC" {
	//	fmt.Println("Part 1 day7.txt result mismatch")
	//}
	//if result != "CABDFE" {
	//	fmt.Println("Part 1 daytest7.txt result mismatch")
	//}
	result := part2(strnarr)
	if result != "EFHLMTKQBWAPGIVXSZJRDUYONC" {
		fmt.Println("Part 1 day7.txt result mismatch")
	}
	if result != "15" {
		fmt.Println("Part 1 daytest7.txt result mismatch")
	}
	//result := part2(strnarr)
	fmt.Println("Day 7 result", result)
}

type Instruction struct {
	FirstStep string
	NextStep  string
}

func parseInputs(inputs []string) (instructions []Instruction) {
	// Step C must be finished before step F can begin.

	for _, s := range inputs {
		var firstStep, nextStep string
		fmt.Sscanf(s, "Step %s must be finished before step %s", &firstStep, &nextStep)
		inst := Instruction{FirstStep: firstStep, NextStep: nextStep}
		instructions = append(instructions, inst)
	}
	return
}

func realSort(input *[]string) {
	SliceUniq(input)
	sort.Strings(*input)
	return
}

func part1(inputs []string) (results string) {
	instructions := parseInputs(inputs)
	fmt.Println(instructions)
	col0 := findFinishedOrder(instructions)
	return strings.Join(col0, "")
}

func SliceUniq(input *[]string) {
	uniq := make(map[string]struct{})

	for _, s := range *input {
		uniq[s] = struct{}{}
	}
	*input = nil
	for k, _ := range uniq {
		*input = append(*input, k)
	}
	return
}
func findNextStep(instructions []Instruction) (nextStep string) {

	var firstFinishedSteps []string
	var nextFinishedSteps []string

	//updatedInstructions := instructions

	firstFinishedSteps = findFirstStepWithNoDependencies(instructions)
	fmt.Println("1 firstFinishedSteps", firstFinishedSteps)
	fmt.Println("")
	firstStepLen := len(firstFinishedSteps)

	nextFinishedSteps = findNextStepWithNoDependencies(instructions)
	//stepOrder = append(stepOrder, nextFinishedSteps...)
	fmt.Println("2 nextFinishedSteps", nextFinishedSteps)
	//fmt.Println("2 stepOrder", stepOrder)
	fmt.Println("")
	nextStepLen := len(nextFinishedSteps)
	if nextStepLen == 0 && firstStepLen == 0 {
		return
	}
	finishedSteps := append(firstFinishedSteps, nextFinishedSteps...)
	fmt.Println("3 finishedSteps\n", finishedSteps)

	realSort(&finishedSteps)
	fmt.Println("4 finishedSteps\n", finishedSteps)

	nextStep = finishedSteps[0]
	return
}

func findFinishedOrder(instructions []Instruction) (stepOrder []string) {

	var firstFinishedSteps []string
	var nextFinishedSteps []string

	//updatedInstructions := instructions

	for i := 0; i < 2000; i++ {
		firstFinishedSteps = findFirstStepWithNoDependencies(instructions)
		fmt.Println("1 firstFinishedSteps", firstFinishedSteps)
		firstStepLen := len(firstFinishedSteps)

		nextFinishedSteps = findNextStepWithNoDependencies(instructions)
		//stepOrder = append(stepOrder, nextFinishedSteps...)
		fmt.Println("2 nextFinishedSteps", nextFinishedSteps)
		//fmt.Println("2 stepOrder", stepOrder)
		nextStepLen := len(nextFinishedSteps)
		if nextStepLen == 0 && firstStepLen == 0 {
			return
		}

		// gather all the finished steps
		finishedSteps := append(firstFinishedSteps, nextFinishedSteps...)
		fmt.Println("3 finishedSteps\n", finishedSteps)

		// sort all the finished steps
		realSort(&finishedSteps)
		fmt.Println("4 finishedSteps\n", finishedSteps)

		// get only the first finished step alphabetically
		finishedSteps = finishedSteps[:1]
		fmt.Println("5 finishedSteps\n", finishedSteps)

		// mark only that first finished step as finishes
		instructions = updateAllSteps(instructions, finishedSteps)
		fmt.Println("3 updatedInstructions\n", instructions)
		// append that first finished step to the step order
		stepOrder = append(stepOrder, finishedSteps...)

	}

	return
}

func findNextStepWithNoDependencies(instructions []Instruction) (finishedSteps []string) {

	notFinished := false
	for _, instruction := range instructions {
		notFinished = false
		possibleFinishedStep := instruction.NextStep
		if possibleFinishedStep == FinishedStep {
			continue
		}
		for _, subInst := range instructions {
			if subInst.NextStep == possibleFinishedStep && subInst.FirstStep != FinishedStep {
				notFinished = true
				break
			}
		}
		if notFinished == false {
			finishedSteps = append(finishedSteps, possibleFinishedStep)
		}
	}
	return
}

func findNextStepWithNoDependenciesPart2(instructions []Instruction) (finishedSteps []string) {

	notFinished := false
	for _, instruction := range instructions {
		notFinished = false
		possibleFinishedStep := instruction.NextStep
		if possibleFinishedStep == FinishedStep {
			continue
		}
		if strings.Contains(possibleFinishedStep, InProgressStep) {
			continue
		}
		for _, subInst := range instructions {
			if subInst.NextStep == possibleFinishedStep && subInst.FirstStep != FinishedStep {
				notFinished = true
				break
			}
		}
		if notFinished == false {
			finishedSteps = append(finishedSteps, possibleFinishedStep)
		}
	}
	return
}

func updateAllSteps(input []Instruction, finishedSteps []string) (ouput []Instruction) {
	for _, uinst := range input {
		for _, finished := range finishedSteps {
			if finished == uinst.FirstStep {
				uinst.FirstStep = FinishedStep
			}
			if finished == uinst.NextStep {
				uinst.NextStep = FinishedStep
			}
		}
		ouput = append(ouput, uinst)
	}
	return
}

func findFirstStepWithNoDependencies(instructions []Instruction) (finishedSteps []string) {
	var notFinished bool
	for _, instruction := range instructions {
		notFinished = false
		possibleFinishedStep := instruction.FirstStep
		if possibleFinishedStep == FinishedStep {
			continue
		}
		for _, subInst := range instructions {
			if subInst.NextStep == possibleFinishedStep {
				notFinished = true
				break
			}
		}
		if notFinished == false {
			finishedSteps = append(finishedSteps, possibleFinishedStep)
		}
	}
	return
}
func findFirstStepWithNoDependenciesPart2(instructions []Instruction) (finishedSteps []string) {
	var notFinished bool
	for _, instruction := range instructions {
		notFinished = false
		possibleFinishedStep := instruction.FirstStep
		if possibleFinishedStep == FinishedStep {
			continue
		}
		if strings.Contains(possibleFinishedStep, InProgressStep) {
			continue
		}
		for _, subInst := range instructions {
			if subInst.NextStep == possibleFinishedStep {
				notFinished = true
				break
			}
		}
		if notFinished == false {
			finishedSteps = append(finishedSteps, possibleFinishedStep)
		}
	}
	return
}

func part2(inputs []string) (results string) {
	instructions := parseInputs(inputs)
	fmt.Println(instructions)
	//nWorkers := 5
	stepNum := doTheWork(instructions)
	//wTime := findWorkTime(instructions, nWorkers)
	return fmt.Sprintf("%v", stepNum)
}

func findWorkTime(instructions []Instruction, nWorkers int) (workTime int) {

	nextStep := findNextStep(instructions)
	for nextStep != "" {

	}

	return
}

func part2psuedo() {

	// workers = me + number of elves
	//nWorkers := 2
	var workers []Worker
	// while there is work left
	for i := 0; i < 2000; i++ {
		//    for each worker
		for _, w := range workers {
			if w.IsWorking {
				continue
			}

			//        is worker busy
			//           yes - continue
			//        is there an unblocked step
			//           no - continue
			//        yes
			//          assign step to worker
			//    end for
		}
		//    is any work done
		//       yes - mark as . in list
		//    increment step count for everyone
		// no work left - we're done here
		return
		//
	}

	// while there is work left
	//    for each worker
	//        status = worker status
	//           yes - continue
	//        is there an unblocked step
	//           no - continue
	//        yes
	//          assign step to worker
	//    end for
	//    is any work done
	//       yes - mark as . in list
	//    increment step count for everyone
	// no work left - we're done here
	//

	// or

	// while there is work left
	//   for each worker
	//    can anyone take work
	//      status = worker.status
	//      if status == waiting
	//        find free step
	//        assign
	//
	//   increment 60 sec
	//
	//   for each aorker
	//     is anyone done
	//        mark step as complete
	//     worker is free
}

type Worker struct {
	AssignedStep  string
	LengthWorking int
	ID            string
	IsWorking     bool
}

//
//func findNextFreeStep(instructions []Instruction) {
//	updatedInstructions := instructions
//	var firstFinishedSteps []string
//	updatedInstructions, firstFinishedSteps = findFirstStepWithNoDependencies(updatedInstructions)
//
//
//}

func findBestStepWithNoDependencies(instructions []Instruction) (aFreeStep string, allFreeSteps []string) {
	// these are all the first steps in the instrctions that are available to work on
	firstStepsWithNoDepenencies := findFirstStepWithNoDependenciesPart2(instructions)
	// these are all the next steps that are available to work on
	nextStepsWithNoDepenencies := findNextStepWithNoDependenciesPart2(instructions)

	// gather all the steps with no dependencies
	stepsWithNoDependencies := append(firstStepsWithNoDepenencies, nextStepsWithNoDepenencies...)
	//fmt.Println("stepsWithNoDependencies\n", stepsWithNoDependencies)

	// sort all the steps with no dependencies
	realSort(&stepsWithNoDependencies)
	//fmt.Printf("\ts %v\n", stepsWithNoDependencies)
	allFreeSteps = stepsWithNoDependencies

	stepWithNoDependencies := ""
	// get only the first steps with no dependencies step alphabetically
	if len(stepsWithNoDependencies) > 0 {
		stepWithNoDependencies = stepsWithNoDependencies[0]
	}
	//fmt.Println("first stepWithNoDependencies\n", stepWithNoDependencies)
	return stepWithNoDependencies, allFreeSteps
}

func getStepTime(step string) (stepTime int) {
	stepTime = SecondsForEachStep
	differenceForThisCharacter := int(step[0] - 'A')
	stepIncrement := differenceForThisCharacter
	stepTime += stepIncrement
	stepTime += 1
	return
}

func printStatus(step int, workers []Worker, finishedSteps []string, freeSteps []string) (stillWorking bool) {

	//Second   Worker 1   Worker 2   Done
	//   0        C          .
	//   1        C          .

	fmt.Printf("\t%v", step)
	// assign if we can
	//assignedToWorker := false
	for _, worker := range workers {
		if worker.IsWorking {
			stillWorking = true
			fmt.Printf("\t%v", worker.AssignedStep)
			fmt.Printf("\t%v", worker.LengthWorking)
			//fmt.Printf("Step:%v Worker %v has been working on step %v for %v seconds\n", step, worker.ID, worker.AssignedStep, worker.LengthWorking)
		} else {
			fmt.Printf("\t%v", ".")

			//fmt.Printf("Step:%v Worker %v is waiting to work\n", step, worker.ID)

		}
		//if worker.IsWorking == false {
		//	stepWithNoDependencies := findBestStepWithNoDependencies(instructions)
		//	if stepWithNoDependencies != "" {
		//		worker.IsWorking = true
		//		worker.AssignedStep = stepWithNoDependencies
		//		worker.LengthWorking = 0
		//		instructions = markStepInProgress(instructions, stepWithNoDependencies)
		//	}
		//}
	}
	//fmt.Printf("StillWorking:%v\n", stillWorking)
	//fmt.Println("finishedSteps", finishedSteps)
	fmt.Printf("\t%v\t%v\n", finishedSteps, freeSteps)

	//fmt.Println()
	return

}
func doTheWork(instructions []Instruction) int {

	var finishedSteps []string
	var workers []Worker
	worker1 := Worker{ID: "ME"}
	worker2 := Worker{ID: "HIM"}
	worker3 := Worker{ID: "HER"}
	worker4 := Worker{ID: "THEM"}
	worker5 := Worker{ID: "SAM"}
	workers = append(workers, worker1)
	workers = append(workers, worker2)
	workers = append(workers, worker3)
	workers = append(workers, worker4)
	workers = append(workers, worker5)
	//stepWithNoDependencies := findBestStepWithNoDependencies(instructions)
	for thisStep := 0; thisStep < 2000; thisStep++ {
		//stepTime := thisStep * SecondsForEachStep
		//printStatus(thisStep,workers, finishedSteps)

		// assign if we can
		//assignedToWorker := false
		var allFreeSteps []string
		for i, _ := range workers {
			if workers[i].IsWorking == false {
				stepWithNoDependencies := ""
				stepWithNoDependencies, allFreeSteps = findBestStepWithNoDependencies(instructions)
				if stepWithNoDependencies != "" {
					workers[i].IsWorking = true
					workers[i].AssignedStep = stepWithNoDependencies
					workers[i].LengthWorking = 0
					instructions = markStepInProgress(instructions, stepWithNoDependencies)
				}
			}
		}
		stillWorking := printStatus(thisStep, workers, finishedSteps, allFreeSteps)
		if !stillWorking {
			return thisStep
		}

		// increment everyone's time
		for i, _ := range workers {
			if workers[i].IsWorking == true {
				workers[i].LengthWorking += 1
				//if worker.LengthWorking == getStepTime(worker.AssignedStep) {
				//	// now our working
				//}
			}
		}
		//printStatus(thisStep, workers, finishedSteps)

		// Find any finished workers
		for i, _ := range workers {
			if workers[i].IsWorking == true {
				if workers[i].LengthWorking >= getStepTime(workers[i].AssignedStep) {
					// now our worker is done
					workers[i].IsWorking = false
					workers[i].LengthWorking = 0
					instructions = markStepFinished(instructions, workers[i].AssignedStep)
					finishedSteps = append(finishedSteps, workers[i].AssignedStep)
					workers[i].AssignedStep = ""
				}
			}
		}

		//fmt.Println(instructions)
		//printStatus(thisStep,workers, finishedSteps)
	}

	//
	//// find workers that are done
	//// increment everyone's time
	//for _, worker := range workers {
	//	if worker.IsWorking == true {
	//		worker.LengthWorking += StepTime
	//	}
	//}
	//
	//
	//
	//
	//var worker Worker
	////-------------------------
	//// ASSIGN this step to a worker
	//// ------------------------
	//
	//if assignedToWorker {
	//	// mark only that first step with no depencicies step as working
	//	worker.IsWorking = true
	//	worker.AssignedStep = stepWithNoDependencies
	//	worker.LengthWorking = 0
	//	instructions = markStepInProgress(instructions, stepWithNoDependencies)
	//}

	return -1
}

func markStepInProgress(input []Instruction, stepWithNoDependencies string) (ouput []Instruction) {
	for _, uinst := range input {
		if stepWithNoDependencies == uinst.FirstStep {
			uinst.FirstStep = stepWithNoDependencies + InProgressStep
		}
		if stepWithNoDependencies == uinst.NextStep {
			uinst.NextStep = stepWithNoDependencies + InProgressStep
		}
		ouput = append(ouput, uinst)
	}
	return
}

func markStepFinished(input []Instruction, stepWithNoDependencies string) (ouput []Instruction) {
	for _, uinst := range input {
		if stepWithNoDependencies+InProgressStep == uinst.FirstStep {
			uinst.FirstStep = FinishedStep
		}
		if stepWithNoDependencies+InProgressStep == uinst.NextStep {
			uinst.NextStep = FinishedStep
		}
		ouput = append(ouput, uinst)
	}
	return
}
