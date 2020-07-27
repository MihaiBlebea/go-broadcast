---
Title: goldmark-meta
Summary: Add YAML metadata to the document
Image: https://yourbasic.org/golang/square-gopher.png
Tags:
    - markdown
    - goldmark
Layout: article
Slug: first-article
Published: "2020-07-28"
---
# Comply Advantage webhook MS

## Problem to be solved

- We want operation team members to do updates in the Comply Advantage platform
- Every update in the platfrom will trigger a webhook to be sent to our V2 API, because of their autosave feature we receive too many payloads instead of one. ex: Marking a search to whitelisted and false_positive will trigger 3-4 payloads to be sent.
- Payload received contains IDs of resources that were changed, but not the full required resource, so for every webhook we receive, we need to trigger a second request to Comply Advantage to receive the whole data before saing it in our database.
- We needed a queue to store these payloads, wait x amount of seconds then send one single paylaod to our V2 monolith

## How to run local?

- Run `make` or `make test-race`
- Or build from the Dockerfile

## How it works?

- We have a map that acts as a queue
- Everytime a webhook payload is received, a new entry is saved in the queue using the `search_id` as a key and `RESEND_INTERVAL_SEC` amount of seconds as ttl
- If we receive a search with the same `search_id` we override the old key and ttl
- A separate go routine checks the queue every `WORKER_SLEEP_INTERVAL_SEC` seconds and if any key is expired then the `search_id` is sent as payload to our V2 monolith to be processed
- If the monolith responds with a different status code then `200`, then a retry counter is incremented and the ttl is also incremented
- If the number of retries to send the `search_id` excedes the `RESEND_MAX_COUNT`, then the key is deleted (worst case scenario to avoid an infinite loop)
- If the `search_id` is successfully sent to V2 monolith, then the key is removed from the queue

#### TODO:
- Add redis as a persistence store 

