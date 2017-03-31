# Practical Promises

=====

```json
{
  "name": "Bojan Djurkovic",
  "location": "Fredericton, NB",
  "work": "Cvent",
  "role": "Lead Software Engineer",
  "facts": [
    "I have an identical twin brother in Ottawa.",
    "Don't worry, he is not at the conference!"
  ]
}
```

NOTES:
_[1 minutes]_

=====

### Thesis

* Promises are not special or revolutionary
* Promises are important for Javascript
* You should use them
* Tips & tricks
* I am not an expert

NOTES:
_[2 minutes]_

- Promises are noting new or revolutionary. Other languages have had them for a while
- Arguably they are not even "good" or "elegant"
- But in my opinion they are important for the Javascript ecosystem
- They are a path to a better future for the language 
- This talk will cover mostly some of the tips and tricks
- I am not an expert. I have not implemented my own Promise library. Just lessons learned from a user.

=====

### Async Javascript

* Callbacks
* Timers
* EventEmitter
* Streams

NOTES:
_[2 minutes]_

- Callbacks are the building blocks of all asynchronous operations in Javascript.
- All other primitives and mechanisms are implemented using callbacks.
- EventEmitters and Streams are only available in Node.js

=====

```js
function handler (params, done) {  
  validate(params, (err, valid) => {
    if (err) return done(err)
    query(valid, (err, docs) => {
      if (err) return done(err)
      performAction(docs, (err, results) => {
        done(err, results)
      })
    })
  })
}
```

NOTES:
_[2 minutes]_

- Traditional Javascript
- "Callback pyramid"
- We can improve using async library. ie. async.waterfall

=====

### What's a Promise?

* Special Javascript object that represents a future value
* Can be either "resolved" or "rejected"
* Have a `then` function that's called when resolved
* Have a `catch` function that's called when rejected

NOTES:
_[1 minutes]_

- Quick overview of what's a Promise

=====

```js
function performActionAsync(docs) {
  return new Promise((resolve, reject) => {
    performAction(docs, (err, results) => {
      if (err) reject(err)
      else resolve(results)
    })
  })
}
```

NOTES:
_[2 minutes]_

- Simple way to implement a promisified function from a callback style function

=====

```js 
function handler (params) {  
  return validate(params)
    .then(query)
    .then(performAction)
    .then((result) => {
      console.log(result)
      return result
    })
}
```

NOTES:
_[2 minutes]_

- Much nicer
- Param automatically follow through

=====

### Questions?
