# budowac-hivemind

AI controller microservice: watches NPC reports from the game server and
publishes high-level intents (goals). Auth-sim lives on `budowac-server`; brains
live here so new species can ship as discrete factory registrations without touching the tick loop.

## Env

| Var | Default | Notes |
| --- | --- | --- |
| `PORT` | `8092` | `/health` |
| `GAME_ID` | `local-dev` | NATS namespace |
| `NATS_URL` | `nats://localhost:4222` | bus |

## NATS

- Sub: `budowac.server.<gameId>.npcs` — NPC report stream
- Pub: `budowac.hivemind.<gameId>.intent` — AI intents (goals)

## Brains

Register new AI in `internal/brain` via `registry.Register(kind, factory)`.
Goat brain: pick a random surface goal when idle or arrived.
