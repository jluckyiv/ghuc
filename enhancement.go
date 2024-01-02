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
// TODO: Use function options instead of With* methods.
func NewEnhancement(baseEnhancement BaseEnhancement) enhancement {
	return enhancement{
		baseEnhancement:      baseEnhancement,
		level:                Level1,
		multipleTarget:       1,
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

func DecrementPrevious(p PreviousEnhancements) PreviousEnhancements {
	// add 4 to avoid negative numbers
	return (p - 1 + 4) % 4
}

func IncrementPrevious(p PreviousEnhancements) PreviousEnhancements {
	return (p + 1) % 4
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

// Enhance* are constants for all the base enhancements.
// They are exported for type safety.
const (
	EnhanceMove BaseEnhancement = iota
	EnhanceJump
	EnhanceAttack
	EnhanceRange
	EnhanceTarget
	EnhanceAddAttackHex
	EnhanceHeal
	EnhanceShield
	EnhanceRetaliate
	EnhanceStrengthen
	EnhanceMuddle
	EnhanceDisarm
	EnhancePierce
	EnhancePoison
	EnhanceWound
	EnhancePush
	EnhancePull
	EnhanceImmobilize
	EnhanceCurse
	EnhanceBless
	EnhanceSpecificElement
	EnhanceAnyElement
	EnhanceSummonsMove
	EnhanceSummonsAttack
	EnhanceSummonsRange
	EnhanceSummonsHP
)

func Title(be BaseEnhancement) string {
	switch be {
	case EnhanceMove:
		return "Move"
	case EnhanceAttack:
		return "Attack"
	case EnhanceRange:
		return "Range"
	case EnhanceShield:
		return "Shield"
	case EnhancePush:
		return "Push"
	case EnhancePull:
		return "Pull"
	case EnhancePierce:
		return "Pierce"
	case EnhanceRetaliate:
		return "Retaliate"
	case EnhanceHeal:
		return "Heal"
	case EnhanceTarget:
		return "Target"
	case EnhancePoison:
		return "Poison"
	case EnhanceWound:
		return "Wound"
	case EnhanceMuddle:
		return "Muddle"
	case EnhanceImmobilize:
		return "Immobilize"
	case EnhanceDisarm:
		return "Disarm"
	case EnhanceCurse:
		return "Curse"
	case EnhanceStrengthen:
		return "Strengthen"
	case EnhanceBless:
		return "Bless"
	case EnhanceJump:
		return "Jump"
	case EnhanceSpecificElement:
		return "Specific Element"
	case EnhanceAnyElement:
		return "Any Element"
	case EnhanceSummonsMove:
		return "Summons Move"
	case EnhanceSummonsAttack:
		return "Summons Attack"
	case EnhanceSummonsRange:
		return "Summons Range"
	case EnhanceSummonsHP:
		return "Summons HP"
	case EnhanceAddAttackHex:
		return "Add Hex"
	default:
		return "Unknown"
	}
}

func Description(be BaseEnhancement) string {
	switch be {
	case EnhanceMove:
		return "adds +1 to move"
	case EnhanceAttack:
		return "adds +1 modifier to attack"
	case EnhanceRange:
		return "adds +1 range to attack"
	case EnhanceTarget:
		return "adds +1 target to attack"
	case EnhanceAddAttackHex:
		return "adds +1 attack hex"
	case EnhanceShield:
		return "adds shield"
	case EnhancePush:
		return "adds push"
	case EnhancePull:
		return "adds pull"
	case EnhancePierce:
		return "adds pierce"
	case EnhanceRetaliate:
		return "adds retaliate"
	case EnhanceHeal:
		return "adds heal"
	case EnhancePoison:
		return "adds poison"
	case EnhanceWound:
		return "adds wound"
	case EnhanceMuddle:
		return "adds muddle"
	case EnhanceImmobilize:
		return "adds immobilize"
	case EnhanceDisarm:
		return "adds disarm"
	case EnhanceCurse:
		return "adds curse"
	case EnhanceStrengthen:
		return "adds strengthen"
	case EnhanceBless:
		return "adds bless"
	case EnhanceJump:
		return "adds jump"
	case EnhanceSpecificElement:
		return "adds specific element"
	case EnhanceAnyElement:
		return "adds any element"
	case EnhanceSummonsMove:
		return "adds +1 move to summons"
	case EnhanceSummonsAttack:
		return "adds +1 modifier to summons attack"
	case EnhanceSummonsRange:
		return "adds +1 range to summons attack"
	case EnhanceSummonsHP:
		return "adds +1 to summons HP"
	default:
		return "Unknown"
	}
}

func ReverseMap[T any](f func(BaseEnhancement) T) map[BaseEnhancement]T {
	return map[BaseEnhancement]T{
		EnhanceMove:            f(EnhanceMove),
		EnhanceAttack:          f(EnhanceAttack),
		EnhanceRange:           f(EnhanceRange),
		EnhanceShield:          f(EnhanceShield),
		EnhancePush:            f(EnhancePush),
		EnhancePull:            f(EnhancePull),
		EnhancePierce:          f(EnhancePierce),
		EnhanceRetaliate:       f(EnhanceRetaliate),
		EnhanceHeal:            f(EnhanceHeal),
		EnhanceTarget:          f(EnhanceTarget),
		EnhancePoison:          f(EnhancePoison),
		EnhanceWound:           f(EnhanceWound),
		EnhanceMuddle:          f(EnhanceMuddle),
		EnhanceImmobilize:      f(EnhanceImmobilize),
		EnhanceDisarm:          f(EnhanceDisarm),
		EnhanceCurse:           f(EnhanceCurse),
		EnhanceStrengthen:      f(EnhanceStrengthen),
		EnhanceBless:           f(EnhanceBless),
		EnhanceJump:            f(EnhanceJump),
		EnhanceSpecificElement: f(EnhanceSpecificElement),
		EnhanceAnyElement:      f(EnhanceAnyElement),
		EnhanceSummonsMove:     f(EnhanceSummonsMove),
		EnhanceSummonsAttack:   f(EnhanceSummonsAttack),
		EnhanceSummonsRange:    f(EnhanceSummonsRange),
		EnhanceSummonsHP:       f(EnhanceSummonsHP),
		EnhanceAddAttackHex:    f(EnhanceAddAttackHex),
	}
}

func Map[T comparable](f func(BaseEnhancement) T) map[T]BaseEnhancement {
	return map[T]BaseEnhancement{
		f(EnhanceMove):            EnhanceMove,
		f(EnhanceAttack):          EnhanceAttack,
		f(EnhanceRange):           EnhanceRange,
		f(EnhanceShield):          EnhanceShield,
		f(EnhancePush):            EnhancePush,
		f(EnhancePull):            EnhancePull,
		f(EnhancePierce):          EnhancePierce,
		f(EnhanceRetaliate):       EnhanceRetaliate,
		f(EnhanceHeal):            EnhanceHeal,
		f(EnhanceTarget):          EnhanceTarget,
		f(EnhancePoison):          EnhancePoison,
		f(EnhanceWound):           EnhanceWound,
		f(EnhanceMuddle):          EnhanceMuddle,
		f(EnhanceImmobilize):      EnhanceImmobilize,
		f(EnhanceDisarm):          EnhanceDisarm,
		f(EnhanceCurse):           EnhanceCurse,
		f(EnhanceStrengthen):      EnhanceStrengthen,
		f(EnhanceBless):           EnhanceBless,
		f(EnhanceJump):            EnhanceJump,
		f(EnhanceSpecificElement): EnhanceSpecificElement,
		f(EnhanceAnyElement):      EnhanceAnyElement,
		f(EnhanceSummonsMove):     EnhanceSummonsMove,
		f(EnhanceSummonsAttack):   EnhanceSummonsAttack,
		f(EnhanceSummonsRange):    EnhanceSummonsRange,
		f(EnhanceSummonsHP):       EnhanceSummonsHP,
		f(EnhanceAddAttackHex):    EnhanceAddAttackHex,
	}
}

func identity(be BaseEnhancement) BaseEnhancement {
	return be
}

func BaseEnhancements() []BaseEnhancement {
	return List(identity)
}

func List[T any](f func(BaseEnhancement) T) []T {
	m := ReverseMap(f)
	list := make([]T, len(m))
	for k, v := range m {
		list[k] = v
	}
	return list
}

// costForBaseEnhancement is a helper function that returns the base cost for
// the base enhancement.
func (e enhancement) costForBaseEnhancement() (Cost, error) {
	var cost Cost
	switch e.baseEnhancement {
	case EnhanceAddAttackHex:
		if e.multipleTarget == 0 {
			return 0, fmt.Errorf("e.multipleTarget is 0")
		}
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
		return Cost(100), nil
	case EnhanceSummonsAttack:
		return Cost(100), nil
	case EnhanceSummonsRange:
		return Cost(50), nil
	case EnhanceSummonsHP:
		return Cost(50), nil
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

// Level* are constants for all the levels, exported for type safety.
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

// costForLevel is a helper function that returns the additional cost for the
// ability card level.
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

// PreviousEnhancements is an enum of all the valid values for previous
// enhancements.
type PreviousEnhancements int

// PreviousEnhancements* are constants for all the valid values for previous
// enhancements, exported for type safety.
const (
	PreviousEnhancements0 PreviousEnhancements = iota
	PreviousEnhancements1
	PreviousEnhancements2
	PreviousEnhancements3
)

// costForPreviousEnhancements is a helper function that returns the
// additional cost for the number of previous enhancements.
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
