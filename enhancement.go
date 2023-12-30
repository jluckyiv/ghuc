package ghec

import "fmt"

// enhancement is a struct that holds the information needed to calculate its
// cost. It is not exported to limit the API surface area. Its only methods are
// With* methods to set its fields and Cost to calculate its cost.
type enhancement struct {
	// baseEnhancement is the base enhancement to calculate the cost.
	// Each base enhancement has a fixed cost.
	baseEnhancement BaseEnhancement
	// level is the level of the ability card to calculate the cost.
	// It must be between 1 and 9.
	level Level
	// multipleTarget serves two purposes:
	// 1. It triggers the multiplier for multiple-target enhancements.
	// 2. It sets the number of current hexes for Add Attack Hex enhancements.
	multipleTarget int
	// previousEnhancements is the number of previous enhancements on the ability
	// card. It must be between 0 and 3.
	previousEnhancements PreviousEnhancements
}

// NewEnhancement creates a new enhancement to calculate its cost.
func NewEnhancement(baseEnhancement BaseEnhancement) enhancement {
	return enhancement{
		baseEnhancement:      baseEnhancement,
		level:                Level1,
		multipleTarget:       0,
		previousEnhancements: PreviousEnhancements0,
	}
}

// WithMultipleTarget sets the number of targets for the enhancement.
// It also sets the number of current hexes for Add Attack Hex enhancements.
func (e enhancement) WithMultipleTarget(multipleTarget int) enhancement {
	e.multipleTarget = multipleTarget
	return e
}

// WithLevel sets the level of the ability card for the enhancement.
func (e enhancement) WithLevel(level Level) enhancement {
	e.level = level
	return e
}

// WithPreviousEnhancements sets the number of previous enhancements on the
// card.
func (e enhancement) WithPreviousEnhancements(previousEnhancements PreviousEnhancements) enhancement {
	e.previousEnhancements = previousEnhancements
	return e
}

// Cost calculates the cost of the enhancement.
// It returns an error if the level or previous enhancements are out of bounds,
// since the With* methods do not validate their inputs.
func (e enhancement) Cost() (Cost, error) {
	if e.level < 1 || e.level > 9 {
		return 0, fmt.Errorf("level must be between 1 and 9, not %d", e.level)
	}
	if e.previousEnhancements < 0 || e.previousEnhancements > 3 {
		return 0, fmt.Errorf("previous enhancements must be between 0 and 3, not %d", e.previousEnhancements)
	}
	baseCost, err := e.costForBaseEnhancement()
	if err != nil {
		return 0, err
	}
	levelCost, err := costForLevel(e.level)
	if err != nil {
		return 0, err
	}
	previousEnhancementCost, err := costForPreviousEnhancements(e.previousEnhancements)
	if err != nil {
		return 0, err
	}
	totalCost := baseCost + levelCost + previousEnhancementCost
	return totalCost, nil
}

// Cost is the cost of an enhancement.
// Probably overkill to have a type for this.
type Cost int

// BaseEnhancement is an enum of all the base enhancements.
type BaseEnhancement int

const (
	EnhanceMove BaseEnhancement = iota
	EnhanceAttack
	EnhanceRange
	EnhanceShield
	EnhancePush
	EnhancePull
	EnhancePierce
	EnhanceRetaliate
	EnhanceHeal
	EnhanceTarget
	EnhancePoison
	EnhanceWound
	EnhanceMuddle
	EnhanceImmobilize
	EnhanceDisarm
	EnhanceCurse
	EnhanceStrengthen
	EnhanceBless
	EnhanceJump
	EnhanceSpecificElement
	EnhanceAnyElement
	EnhanceSummonsMove
	EnhanceSummonsAttack
	EnhanceSummonsRange
	EnhanceSummonsHP
	EnhanceAddAttackHex
)

func (e enhancement) costForBaseEnhancement() (Cost, error) {
	var cost Cost
	switch e.baseEnhancement {
	case EnhanceAddAttackHex:
		return Cost(200 / e.multipleTarget), nil
	case EnhanceMove:
		cost = 30
	case EnhanceAttack:
		cost = 50
	case EnhanceRange:
		cost = 30
	case EnhanceShield:
		cost = 100
	case EnhancePush:
		cost = 30
	case EnhancePull:
		cost = 30
	case EnhancePierce:
		cost = 30
	case EnhanceRetaliate:
		cost = 100
	case EnhanceHeal:
		cost = 30
	case EnhanceTarget:
		cost = 50
	case EnhancePoison:
		cost = 75
	case EnhanceWound:
		cost = 75
	case EnhanceMuddle:
		cost = 50
	case EnhanceImmobilize:
		cost = 100
	case EnhanceDisarm:
		cost = 150
	case EnhanceCurse:
		cost = 75
	case EnhanceStrengthen:
		cost = 50
	case EnhanceBless:
		cost = 50
	case EnhanceJump:
		cost = 50
	case EnhanceSpecificElement:
		cost = 100
	case EnhanceAnyElement:
		cost = 150
	case EnhanceSummonsMove:
		cost = 100
	case EnhanceSummonsAttack:
		cost = 100
	case EnhanceSummonsRange:
		cost = 50
	case EnhanceSummonsHP:
		cost = 50
	default:
		return 0, fmt.Errorf("unknown base enhancement %d", e.baseEnhancement)
	}
	if e.multipleTarget > 1 {
		cost *= 2
	}
	return cost, nil
}

// Level is an enum of all the levels.
// Probably overkill to have an enum for this.
type Level int

const (
	Level1 Level = 1
	Level2 Level = 2
	Level3 Level = 3
	Level4 Level = 4
	Level5 Level = 5
	Level6 Level = 6
	Level7 Level = 7
	Level8 Level = 8
	Level9 Level = 9
)

func costForLevel(level Level) (Cost, error) {
	switch level {
	case Level1:
		return 0, nil
	case Level2:
		return 25, nil
	case Level3:
		return 50, nil
	case Level4:
		return 75, nil
	case Level5:
		return 100, nil
	case Level6:
		return 125, nil
	case Level7:
		return 150, nil
	case Level8:
		return 175, nil
	case Level9:
		return 200, nil
	default:
		return 0, fmt.Errorf("level must be between 1 and 9, not %d", level)
	}
}

type PreviousEnhancements int

const (
	PreviousEnhancements0 PreviousEnhancements = iota
	PreviousEnhancements1
	PreviousEnhancements2
	PreviousEnhancements3
)

func costForPreviousEnhancements(previousEnhancements PreviousEnhancements) (Cost, error) {
	switch previousEnhancements {
	case PreviousEnhancements0:
		return 0, nil
	case PreviousEnhancements1:
		return 75, nil
	case PreviousEnhancements2:
		return 150, nil
	case PreviousEnhancements3:
		return 225, nil
	default:
		return 0, fmt.Errorf("previous enhancements must be between 0 and 3, not %d", previousEnhancements)
	}
}
