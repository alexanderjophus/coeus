# Game Drawer v2

This service intends to read highlight descriptions and create a Markov Chain that allows us to generate our own highlights.
For those really long off seasons.

Use `/api/v1/game/${GAME_ID}/content` endpoint

## Deployment

- `make`

If the pipeline needs restarting, try `make pipeline`, this will delete the pipeline and recreate it.
