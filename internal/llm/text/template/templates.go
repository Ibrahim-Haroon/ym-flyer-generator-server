package template

import "fmt"

const Role = `
You're someone who creates background templates for flyers by describing in great details (max 900 chars)
The design will then be given verbatim to a image generation model to create.
`

func ImageDescriptonGenerationPrompt(colorPalette string) string {
	return fmt.Sprintf(`
		You need to write the description to create a background template for a flyer that will be used to promote
        a YM (young muslims) event. The background *must* be empty in the center because text and logos will be
        overlaid on top later by a human. It is crucial that the design is symmetrical.

        The color palette should ONLY include %s.

        Here are some perfect examples to follow:
        Example 1:
        Islamic-themed border design in portrait orientation. Features a large central white space (80%% of design)
        with delicate watercolor frame in soft blue-grey tones. Top border has alternating crescents connected by
        thin lines, with hanging lanterns (one central yellow, others grey-blue). Side borders are minimal with
        light watercolor wash. Bottom shows matching minarets with detailed towers and spires, plus three domed
        structures with golden finials between them. Overall watercolor style creates ethereal effect with subtle
        water spots in corners. Design maintains perfect symmetry. Central empty space allows for text/content
        addition, taking up roughly 60%% height and 50%% width of total design.

        Example 2:
        Islamic-themed border design in portrait orientation. Large empty white/cream space occupies central 80%% of
        design (from 15%% to 85%% height, 20%% to 80%% width), this space MUST remain completely blank for content.
        Decorative elements confined strictly to borders: Top 15%% has thin arch with subtle geometric pattern, small
        golden crescent moon. Bottom 15%% shows minimalist mosque silhouette (contained entirely within bottom
        border). Side borders (each 10%% of total width) are simple geometric patterns in gray with gold accents. No
        patterns or watermarks in content area, design also maintains perfect symmetry. Overall palette: grays,
        white/cream, touches of gold.
	`, colorPalette)
}
