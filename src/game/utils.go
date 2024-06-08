package game

func ApplyStatsValuesFromNames(caster Entity, target Entity, stats []*Stat) {
	for _, stat := range stats {
		switch stat.Name { 
		case "Health":
			target.TakeDamage(caster, -stat.Value)
			break
		case "Damage":
			target.TakeDamage(caster, stat.Value)
			break
		}
	}
}

func SkillFromName(name string) *Skill {
	for _, skill := range Skills {
		if skill.Name == name {
			return skill
		}
	}
	return nil
}
