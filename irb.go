package irb

import (
	"math"
	"errors"
	"fmt"
)

type Result float64

const (
	Win   = Result(1)
	Draw  = Result(0.5)
	Loose = Result(0)
)

type IRB struct {
	RA               float64
	RB               float64
	SA               Result
	SB               Result
	IsBigWin         bool
	Weighting        float64
	BigWinMultiplier float64
	RAHandicapped    float64
	RBHandicapped    float64
	RatingGap        float64
	CoreRatingChange float64
	RAN              float64
	RBN              float64
}

const (
	homeHandicap            = float64(3)
	defaultWeighting        = float64(1)
	defaultBigWinMultiplier = float64(1.5)
)

func New(home, away float64, sa, sb Result) (*IRB, error) {

	if float64(sa)+float64(sb) != 1 {
		fmt.Println("hello")
		return nil, errors.New("invalid result")
	}

	var i = &IRB{
		RA:               home,
		RB:               away,
		SA:               sa,
		SB:               sb,
		IsBigWin:         false,
		Weighting:        defaultWeighting,
		BigWinMultiplier: defaultBigWinMultiplier,
	}

	i.calculate()

	fmt.Println(i)

	return i, nil
}

func (i *IRB) calculate() {
	i.allowForHomeAdvantage()
	i.calculateTheRatingGap()
	i.checkThePossibleCoreRatingChanges()
	i.applyWeightingFactors()
}
func (i *IRB) allowForHomeAdvantage() {
	i.RAHandicapped = i.RA + homeHandicap
	i.RBHandicapped = i.RB
}
func (i *IRB) calculateTheRatingGap() {
	i.RatingGap = math.Abs(i.RAHandicapped - i.RBHandicapped)
}
func (i *IRB) checkThePossibleCoreRatingChanges() {

	if i.RAHandicapped > i.RBHandicapped && i.SA == Win {
		i.CoreRatingChange = 1 - (i.RatingGap / 10)
	}

	if i.RAHandicapped > i.RBHandicapped && i.SB == Win {
		i.CoreRatingChange = 1 + (i.RatingGap / 10)
	}

	if i.RAHandicapped < i.RBHandicapped && i.SA == Win {
		i.CoreRatingChange = 1 + (i.RatingGap / 10)
	}

	if i.RAHandicapped > i.RBHandicapped && i.SB == Win {
		i.CoreRatingChange = 1 - (i.RatingGap / 10)
	}

	if i.RAHandicapped > i.RBHandicapped && i.SA == Draw && i.SB == Draw {
		i.CoreRatingChange = i.RatingGap / 10
	}
}
func (i *IRB) applyWeightingFactors() {
	i.CoreRatingChange *= i.Weighting
	if i.IsBigWin {
		i.CoreRatingChange *= i.BigWinMultiplier
	}

	if i.RAHandicapped > i.RBHandicapped && i.SA == Win {
		i.RAN = i.RA + i.CoreRatingChange
		i.RBN = i.RB - i.CoreRatingChange
	}

	if i.RAHandicapped > i.RBHandicapped && i.SB == Win {
		i.RAN = i.RA - i.CoreRatingChange
		i.RBN = i.RB + i.CoreRatingChange
	}

	if i.RAHandicapped < i.RBHandicapped && i.SA == Win {
		i.RAN = i.RA - i.CoreRatingChange
		i.RBN = i.RB + i.CoreRatingChange
	}

	if i.RAHandicapped < i.RBHandicapped && i.SB == Win {
		i.RAN = i.RA + i.CoreRatingChange
		i.RBN = i.RB - i.CoreRatingChange
	}

	if i.RAHandicapped > i.RBHandicapped && i.SA == Draw && i.SB == Draw {
		i.RAN = i.RA - i.CoreRatingChange
		i.RBN = i.RB + i.CoreRatingChange
	}
}
