# API Walk

## Dependencies

- Router [chi](https://github.com/go-chi/chi)

## Routes

- __GET__ `/` - basic output
- __POST__ `/user` - register user
- __DELETE__ `/user/:userId` - delete user
- __GET__ `/:userId/walk` - get user's walk sessions
- __POST__ `/:userId/walk/start` - start a walk
- __POST__ `/:userId/walk/stop` - stop the current walk
- __POST__ `/:userId/walk/:walkId` - put coordinates for a current walk
- __GET__ `/:userId/walk/:walkId` - get route of a walk as coordinates list
- __DELETE__ `/:userId/walk/:walkId` - delete a walk route