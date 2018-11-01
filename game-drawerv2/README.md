# Game Drawer v2

This service draws a heatmap of a games activities.
It takes in a single arg for the type of event we're plotting.

Use `/api/v1/game/${GAME_ID}/feed/live` endpoint

## Deployment

- `make`

If the pipeline needs restarting, try `make pipeline`, this will delete the pipeline and recreate it.
