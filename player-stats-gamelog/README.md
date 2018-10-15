# Star Tracker

This service is designed to read `/api/v1/people/${PEOPLE_ID}/stats?stats=gameLog` and creates some pretty pictures

## Deployment

- `make`

If the pipeline needs restarting, try `make pipeline`, this will delete the pipeline and recreate it.
