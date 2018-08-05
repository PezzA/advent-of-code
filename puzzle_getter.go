package main

import (
	"errors"

	"github.com/pezza/AoC2017/2015/Day201501"
	"github.com/pezza/AoC2017/2015/Day201502"
	"github.com/pezza/AoC2017/2015/Day201503"
	"github.com/pezza/AoC2017/2015/Day201504"
	"github.com/pezza/AoC2017/2015/Day201505"
	"github.com/pezza/AoC2017/2015/Day201506"
	"github.com/pezza/AoC2017/2015/Day201508"
	"github.com/pezza/AoC2017/2015/Day201524"
	"github.com/pezza/AoC2017/2015/Day201525"
	"github.com/pezza/AoC2017/2016/Day201611"
	"github.com/pezza/AoC2017/2016/Day201617"
	"github.com/pezza/AoC2017/2016/Day201619"
	"github.com/pezza/AoC2017/2016/Day201622"
	"github.com/pezza/AoC2017/2017/Day201701"
	"github.com/pezza/AoC2017/2017/Day201702"
	"github.com/pezza/AoC2017/2017/Day201703"
	"github.com/pezza/AoC2017/2017/Day201704"
	"github.com/pezza/AoC2017/2017/Day201705"
	"github.com/pezza/AoC2017/2017/Day201706"
	"github.com/pezza/AoC2017/2017/Day201707"
	"github.com/pezza/AoC2017/2017/Day201708"
	"github.com/pezza/AoC2017/2017/Day201709"
	"github.com/pezza/AoC2017/2017/Day201710"
	"github.com/pezza/AoC2017/2017/Day201711"
	"github.com/pezza/AoC2017/2017/Day201712"
	"github.com/pezza/AoC2017/2017/Day201713"
	"github.com/pezza/AoC2017/2017/Day201714"
	"github.com/pezza/AoC2017/2017/Day201715"
	"github.com/pezza/AoC2017/2017/Day201716"
	"github.com/pezza/AoC2017/2017/Day201717"
	"github.com/pezza/AoC2017/2017/Day201718"
	"github.com/pezza/AoC2017/2017/Day201719"
	"github.com/pezza/AoC2017/2017/Day201720"
	"github.com/pezza/AoC2017/2017/Day201721"
	"github.com/pezza/AoC2017/2017/Day201722"
	"github.com/pezza/AoC2017/2017/Day201723"
	"github.com/pezza/AoC2017/2017/Day201724"
	"github.com/pezza/AoC2017/2017/Day201725"
	"github.com/pezza/AoC2017/TestDay"
	"github.com/pezza/aoc2017/2015/Day201507"
	"github.com/pezza/aoc2017/2015/Day201510"
)

func getPuzzle(day int, year int) (dailyPuzzle, error) {
	dailyPuzzles := [...]dailyPuzzle{
		dayEntry.Entry,
		Day201501.Entry,
		Day201502.Entry,
		Day201503.Entry,
		Day201504.Entry,
		Day201505.Entry,
		Day201506.Entry,
		Day201507.Entry,
		Day201508.Entry,
		Day201510.Entry,
		Day201524.Entry,
		Day201525.Entry,
		Day201611.Entry,
		Day201617.Entry,
		Day201619.Entry,
		Day201622.Entry,
		Day201701.Entry,
		Day201702.Entry,
		Day201703.Entry,
		Day201704.Entry,
		Day201705.Entry,
		Day201706.Entry,
		Day201707.Entry,
		Day201708.Entry,
		Day201709.Entry,
		Day201710.Entry,
		Day201711.Entry,
		Day201712.Entry,
		Day201713.Entry,
		Day201714.Entry,
		Day201715.Entry,
		Day201716.Entry,
		Day201717.Entry,
		Day201718.Entry,
		Day201719.Entry,
		Day201720.Entry,
		Day201721.Entry,
		Day201722.Entry,
		Day201723.Entry,
		Day201724.Entry,
		Day201725.Entry,
	}

	for _, puzzle := range dailyPuzzles {
		puzzleYear, puzzleDay, _ := puzzle.Describe()
		if puzzleDay == day && puzzleYear == year {
			return puzzle, nil
		}
	}

	return nil, errors.New("day specified has not been fully implemented yet")
}
