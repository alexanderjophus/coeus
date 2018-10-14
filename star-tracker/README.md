# Star Tracker

This service is designed to read `/api/v1/game/${GAME_ID}/feed/live` data and exports the star player data

## Deployment

(tidy this into a build/make step)

- `make`

If the pipeline needs restarting, try `make pipeline`, this will delete the pipeline and recreate it.
