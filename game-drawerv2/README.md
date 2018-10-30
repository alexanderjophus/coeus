# Play Tracker

DEPRECATED in favour of game simplifier

This service creates a flat csv of all the plays that happened in the game.

Use `/api/v1/game/${GAME_ID}/feed/live` endpoint

## Deployment

- `make`

If the pipeline needs restarting, try `make pipeline`, this will delete the pipeline and recreate it.
