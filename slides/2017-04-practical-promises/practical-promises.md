# Practical 
# Promises

=====

### package.json

```json
{
  "name": "Bojan Djurkovic",
  "config": {
    "location": "Fredericton, NB",
    "work": "Cvent",
    "role": "Lead Software Engineer",
  },
  "devDependencies": {
    "osx": "^10.11.0",
    "vscode": "^1.10.0"
  },
  "dependencies": {
    "food": "latest",
    "water": "^1.0.0"
  }
}
```

NOTES:
_[1 minutes]_

=====

### Thesis

* Promises are not really special or revolutionary
* Promises are important and beneficial for Javascript
* You should use them
* Tips, tricks & lots of code
* I am not an expert

NOTES:
_[2 minutes]_

- Promises are nothing new or revolutionary. Other languages have had them for a while.
- But in my opinion they are important for the Javascript ecosystem
- They are a path to a better future for the language 
- This talk will cover mostly some of the tips and tricks and code
- I am not an expert. I have not implemented my own Promise library. Just lessons learned from a user.

=====

### Why this talk? Why now?

* Promises have been in browser and Node for 2 years
* Most devs still prefer callbacks
* Modules with Promise based API's are still an exception

NOTES:
_[2 minutes]_

* Node core API is still Prmose-based and maps nicely to underlying v8 platform
* `XHR` is callback based, and `fetch()` adoption lagging
* Bluebird and Q still popular Promise libraries, which just leads to fragmantation
* Also async / await only works with native Promises
* Express still order of magnitude more popular than Koa, the Promise-based spiritual successor to Express.
* Technical and comminity issue

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

### Traditinal Callback approach

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
- Resolved also called "fulfilled"
- `then()` can optionally hava a rejection handler

=====

### Benefits

* More radable code
* Forces consistent asynchronicity
* Zalgo-safe
* Callbacks can be called multiple times, Promise are resolved once
* Allows for simpler, more efficient memoization implementations and patterns

NOTES:
_[3 minutes]_

* Even though promises are usually ‘future’ data, once we actually have a promise we don’t need to care whether the data will be there in future, or it’s already been resolved. We call `then()` in either case. As such, promises force consistent asynchronicity 
* Zalgo-sage - we do not need to care whther we are actually on the same tick or not
* Simple memoization, we just need to cache the promise

=====

### Creating a Promise

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
- Call resolve with fulfilment value
- Call reject with error / reason for rejection

=====

### An improvement

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
- `then` method returns a `Promise` which allows for method chaining
- Params automatically follow through

=====

### CHAIN, DON'T NEST

```js
function handler (params) {  
  return new Promise(resolve, reject {
    validate(params)
      .then(params => {
        performAction(params)
          .then(result => {
            console.log(result)
            resolve(result)  
          })
      })
  })
}
```

NOTES:
_[2 minutes]_

=====

### Promise.resolve()

* Use `Promise.resolve` to turn any value into a `Promise`

```js
const p1 = Promise.resolve('foo')
p1.then(str => {
  console.log(str)
})
```

NOTES:
_[1 minutes]_

=====

### Promise.reject()

```js
Promise.reject(new Error('Boom!'))
  .then(str => {
    console.log('should not be here %s', str)
  }, err => {
    console.error(err)
  })
```

```sh
Error: Boom!
    at Object.<anonymous> (/app.js:1:89)
    at Module._compile (module.js:571:32)
    at Object.Module._extensions..js (module.js:580:10)
    at Module.load (module.js:488:32)
```

* `then` can take the rejection handler

NOTES:
_[1 minutes]_

=====

### Another example

```js
callAPI()
  .then(
    handleResponse, 
    handleError
  )
```

* What happens if `handleResponse` or `handleError` crash?

=====

### Always use `catch()`

```js
callAPI()
  .then(
    handleResponse, 
    handleAPIError
  )
  .catch(handleProgrammerOrSystemError)
```

* `then` error handler is optional

=====

### Promise.all() for parallel tasks

```js
Promise.all([
  task1(),
  task2(),
  task3()
])
  .then(vals => {
    console.log(vals)  // ['foo', 'bar', 42]
  })
  .catch(e => {
    console.error('One of the tasks failed!', e)
  })
```

* `Promise.all([ ... ])` waits for all fulfillments (or the first rejection)

NOTES:
_[1 minutes]_

- Tasks are promise returning functions
- Note that you can pass just a primitive value to the array and it will be resolved

=====

### Promise.race()

```js
Promise.race([
  task1(),
  task2(),
  task3(),
  42
])
  .then(result => {
    console.log(result)  // 42
  })
  .catch(e => {
    console.error('One of the tasks failed!', e)
  })
```

* `Promise.race([ ... ])` waits only for either the first fulfillment or rejection.

NOTES:
_[1 minutes]_

=====

### Breaking a promise chain ?

```js
doTask()
	.then(alwaysRun1())
	.then(condition => condition || somehowBreakTheChain())
	.then(onlyRunConditional1())
	.then(onlyRunConditional2())
	.then(alwaysRun2())
```

=====

### Solution

```js
function runConditional(conditional) {
  Promise.resolve(conditional)
    .then(onlyRunConditional1())
    .then(onlyRunConditional2())
}

doTask()
	.then(alwaysRun1())
	.then(condition => condition && runConditional(condition))
	.then(alwaysRun2())
```

=====

### Promisifying

* Converting callback style functions into Promise returning ones
* Many utilities exist; I like `pify`

```js
const fs = require('fs')
const pify = require('pify')

// promisify a single function
const readFileAsync = pify(fs.readFile)
readFileAsync('package.json', 'utf8').then(data => {
	console.log(JSON.parse(data).name)  // 'pify'
})

// or promisify all functions in a module
const fsAsync = pify(fs)
fsAsync.readFile('package.json', 'utf8').then(data => {
	console.log(JSON.parse(data).name)  // 'pify'
})
```

=====

### Providing an API: Options

* Do nothing and let client convert it how they want to
* Provide a Promisified version of module separately. ie. `foo-lib-async`
* Separate callback and Promisified functions within a single module `foo-lib`
* Single API that determines the mechanism at run time based on params
  - If callback provided return via callback
  - If no callback, return Promise

NOTES:
_[2 minutes]_

- I prefer last approach

=====

### Separate API

```js
const pify = require('pify')

function foo (fn) {
  // ... callback-style function
}

module.exports.foo = foo
module.exports.fooAsync = pify(foo)
```

NOTES:
_[1 minutes]_

* It's become a somehwat of custom to name the Promised-based functions in these cases with "Async" suffix

=====

### Single API

```js
const pifyCall = require('promisify-call')

function foo (fn) {
  // ... callback-style function
}

module.exports = {
  foo: function () {
    pifyCall(this, foo, ...arguments)
  }
}
```

NOTES:
_[1 minutes]_

=====

### Single API 2

```js
const nodeify = require('promise-nodeify')

function foo () {
  // ... Promise-returning function
}

module.exports = {
  foo: function (fn) {
    return nodeify(foo(), fn)
    // if callback is not a function, promise is returned as-is
    // otherwise callback will be called when 
    // promise is resolved or rejected
  }
}
```

NOTES:
_[1 minutes]_

=====

### Dealing with concurrency

* `Promise.all([...])` executes all promises at once. How can we limit concurrency?

```js
const pAll = require('p-all')

const tasts = [
  () => task1(),
  () => task2(),
  () => task3(),
  () => task4(),
  () => task5()
]

pAll(tasts, {concurrency: 2})
  .then(result => {
	  console.log(result)
  })
```

NOTES:
_[1 minutes]_

=====

### Mapping

```js
const pMap = require('p-map')
const got = require('got')

const sites = [
  getUrl(), // return a Promise
  'github.com',
  'standardjs.com',
  'nodejs.org'
]

const get = url => got(url).then(res => res.statusMessage)

pMap(sites, get, {concurrency: 2}).then(results => {
  console.log(results) // [ 'OK', 'OK', 'OK', 'OK' ]
})
```

NOTES:
_[1 minutes]_

=====

### Use Native Promises

* Native Promises are awesome for compatibility and future-proofness
* Libraries like `bluebird` and `Q` add extras
* There is a whole ecosystem of tiny functional modules for native Promises
* https://github.com/wbinnssmith/awesome-promises
* In the real world Promise implementation is unlikely your performance bottleneck
* Native Promises will only get faster
* Async / await returns a native Promise no matter what implementation you use

NOTES:
_[3 minutes]_

=====

### Async / Await

* The real exciting and _practical_ evolution of async Javascript
* An `async` function returns a Promise
* Use `await` on a Promise to wait for it be resolved or an exception thrown
* Just act on Promises under the hood

NOTES:
_[1 minutes]_

- Promises are just the necessary, not as elegant, stepping stone towards a better mechanism

=====

### Example

```js
const got = require('got')

async function main () {
  try {
    const res = await got('github.com')
    console.log(res.statusMessage)
  } catch (e) {
    console.error(e)
  }
}

main()  // OK
```

* Executes like it reads!

NOTES:
_[1 minutes]_

=====

### Pitfalls

* Not using `await`

```js
let res = got('github.com') // res is the Promise
res = await got('github.com') // res is the response
```

* awaiting multiple values

```js
let foo = await getFoo()
let bar = await getBar()

let [foo2, bar2] = await Promise.all([getFoo(), getBar()])
```

* Not handling errors; use normal `try` and `catch`

=====

### Sequential async

```js
async function get(url) {
  const r = await got(url)
  return r.statusMessage
}

async function main () {
  const urls = ['github.com', 'standardjs.com','nodejs.org']
  for (let url of urls) {
    const msg = await get(url)
    console.log(msg)
  }
}

main()
// OK
// OK
// OK
```

NOTES:
_[1 minutes]_

=====

### Concurrent execution

```js
const got = require('got')

async function get(url) {
  const r = await got(url)
  return r.statusMessage
}

async function main () {
  const urls = ['github.com', 'standardjs.com','nodejs.org']
  const promises = urls.map(url => get(url))
  let results = await Promise.all(promises);
  console.log(results); // [ 'OK', 'OK', 'OK' ]
}

main() 
```

NOTES:
_[1 minutes]_

=====

### Review - callbacks

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
_[1 minutes]_

=====

### Review - Promises

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
_[1 minutes]_

=====

### Review - Async / Await

```js
async function handler (params) {  
  const valid = await validate(params)
  const docs = await query(valid)
  return performAction(docs)
}
```

NOTES:
_[1 minutes]_

=====

### Final thoughts

* With Promises a much more practical and elegant async mechanism exists in async /await
* async / awiait is currently supported in Node.js LTS and most browsers
* Native Promise-based Node.js is coming eventually
* Start using Promises!

=====

### Questions?
