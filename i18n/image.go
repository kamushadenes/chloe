package i18n

import "math/rand"

// TODO: add actual localization support, for now this only holds some english messages

func GetImageGenerationText() string {
	var messages = []string{
		"Hold on, I'll attempt to generate this picture.",
		"Just a moment, let me work on producing this visual.",
		"Wait a sec, I'm going to endeavor to form this image.",
		"Bear with me, I'll strive to bring this picture to life.",
		"Be patient, I'm making an effort to design this graphic.",
		"Stay with me, I'll give it a shot at crafting this illustration.",
		"One moment, please, as I attempt to construct this visual representation.",
		"Give me a minute, I'll do my best to fabricate this image.",
		"Allow me some time, I'm going to try and render this picture.",
		"Kindly wait, I'll put in the effort to create this depiction.",
		"I'll try to create that.",
	}

	return messages[rand.Intn(len(messages))]
}
