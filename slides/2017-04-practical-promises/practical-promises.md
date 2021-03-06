## **PRACTICAL**
## **PROMISES**

NOTES:
_[1 minutes]_

- Thank the organizers and the audience

=====

```json
{
  "name": "Bojan Djurkovic",
  "config": {
    "location": "Fredericton, NB",
    "work": "Cvent",
    "role": "Lead Software Engineer",
    "twitter": "bojantweets",
    "github": "bojand"
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

- Hello, my name is Bojan Djurkovic.
- I am a Lead Software Engineer for Cvent. It's a company based out of Washington DC.
- Coincidentally we make software for conferences, conventions and events not unlike this one.
- I work mostly on server side with Node.js, some React, and unfortunately some Java here and there.

=====

### Thesis

* Promises are not really revolutionary
* Promises are important and beneficial for Javascript
* You should use them
* Tips, tricks & lots of code

NOTES:
_[2 minutes]_

- Promises are nothing new or revolutionary. Other languages have had them for a while. 
- But in my opinion they are important for the Javascript ecosystem
- They are a path to a better future for the language 
- You should use them
- This talk will cover mostly tips and tricks and include lots of code
- I would not consider myself an expert or authority. I have not implemented my own Promise library. These are just lessons learned from a user.

=====

### Why this talk? Why now?

* Promises have been in browser and Node for 2 years
* Most developers still prefer callbacks
* Modules with Promise based API's are still an exception

NOTES:
_[2 minutes]_

- Promises have been in browser and Node for 2 years, yet I think there still seems to be this aversion or at least lack of will to adopt them and use them.
- Bluebird and Q still popular Promise libraries, which just leads to fragmantation, which will only cause problems in future
- Express still order of magnitude more popular than Koa, the Promise-based spiritual successor to Express.
- This is due partly to technical reasons. Node core API is callback-based because maps nicely to underlying v8 platform, and so since it's the root API that everyone has to use, the paradigm carries through into the user land.
- But Promise-based Node will come eventually.
- Similarly `XHR` and ajax is event / callback based, and `fetch()` adoption lagging
- Technical and comminity issues

=====

### Traditional Async Javascript

* Callbacks
* Timers
* EventEmitter
* Streams

NOTES:
_[1 minutes]_

- What are the different mechanisms for performing asynchronous actions in Javascript?
- Callbacks are the building blocks of all asynchronous operations in Javascript.
- All other primitives and mechanisms, such as Timers, are implemented using callbacks.
- EventEmitters and Streams are only available in Node.js

=====

### Traditional callback approach

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
- Duplicate error handling code
- We can improve using async library. ie. `async.waterfall`

=====

### What's a Promise?

* Special Javascript object that represents a future value
* Can be either "resolved" or "rejected"
* Have a `then` function that's called when resolved
* Have a `catch` function that's called when rejected

NOTES:
_[2 minutes]_

- Just a quick overview of what's a Promise. It's a special Javascript object that represents some future value of computation.
- Promises can be resolved (or fulfilled) when the async action succeeds
- Or it can be "rejected" when the action fails
- A Promise has a `then` function that's executed when the promise is resolved
- `then()` must take a function that returns either a Promise, a value (including implicid `undefined`) or throw.
- `then()` can optionally hava a rejection handler
- A Promise has a `catch` function that's called when rejected

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

performActionAsync(someDocs).then(
  result => console.log(result), 
  err => console.error(err)
)
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

- If we take our example code from earlier, and implement it using Promises and it looks much nicer
- `then` method returns a `Promise` which allows for method chaining
- Parameters automatically flow through
- `async.waterfall` for free

=====

### Some benefits

* More readable code
* Consistent async idioms
* Zalgo-safe
* Callbacks can be called multiple times, Promise are resolved once
* Allows for simpler, more efficient memoization implementations and patterns

NOTES:
_[3 minutes]_

- Usage of Promises can result in more readable code.
- Your code will become smaller, more elegant, and easier to reason about. 
- Promises force us to use consistent async idioms. Even though promises are usually ‘future’ data, once we actually have a promise we don’t need to care whether the data will be there in future, or it’s already been resolved. We call `then()` in either case.
- Zalgo-safe - we do not need to care whther we are actually on the same tick or not
- Callbacks can be called multiple times, Promise are resolved once
- Simple memoization, we just need to cache the promise

=====

### Chain, don't nest

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
_[1 minutes]_

- This is effectively the same code, but it defeats the purpose

=====

### `Promise.resolve()`

* Use `Promise.resolve` to turn any value into a `Promise`

```js
const p1 = Promise.resolve('foo')
p1.then(str => {
  console.log(str)
})
```

NOTES:
_[1 minutes]_

- Useful for wrapping anything into a Promise

=====

### `Promise.reject()`

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

- `reject()` does not need to take an Error

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


NOTES:
_[1 minutes]_

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

NOTES:
_[1 minutes]_

=====

### `Promise.all()` for parallel tasks

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
- `Promise.all([ ... ])` waits for all fulfillments (or the first rejection)
- Note that you can pass just a primitive value to the array and it will be resolved

=====

### `Promise.race()`

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

- `Promise.race([ ... ])` waits only for either the first fulfillment or rejection.

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

NOTES:
_[1 minutes]_

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

NOTES:
_[1 minutes]_

- Here we use the conditional to resolve on it and nest the rest of the conditional logic behind it

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

NOTES:
_[1 minutes]_

- Promisifying is the mechanism of converting callback style functions into Promise returning ones
- There are many utilities and libraries that already exist to help with this goal
- I prefer `pify` it's nice and simple
- We can promisify just a single function, or a whole module

=====

<section>
  <h3>Providing an API: Options</h3>
  <ul>
    <li class="fragment fade-in">Do nothing</li>
    <li class="fragment fade-in">Provide a promisified version of module separately. ie. `foo-lib-async`</li>
    <li class="fragment fade-in">Separate callback and promisified API within a single module `foo-lib`</li>
    <li class="fragment fade-in">Single API that determines the mechanism at run time based on invocation
      <ul>
        <li>If callback provided return via callback</li>
        <li>If no callback, return Promise</li>
      <ul>
    </li>
  </ul>

<aside class="notes">
NOTES:
_[3 minutes]_

- What are our options for providing a Promise based API for our module or library?
- Of course we can do nothing. Just provide a callback based API. There is certainly nothing wrong with that. And the client can take the approach they prefer.
- We can provide a complete separate library that exposes a Promise-based API only. For example our `foo-lib-async` or something.
- We can provide both callback and Promise-based API's within the single module.
- Or we can provide a single API that determines the mechanism at run time based on invocation
  - If callback provided return via callback
  - If no callback, return Promise
- I prefer this last approach, but that's just a personal preference. I don't believe any single approach is necessarily better than the other.
</aside>
</section>

=====

### Separate API in a single module

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

* It has become a somewhat of custom to name the Promised-based functions in these cases with "Async" suffix

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

- This is an example where we have a callback based implementation and we're converting it to allow both callbacks and Promises.
- One negative of this single API approach is that the client has to explicitly set the callback even in cases where they do not care about the result. For example they have to explicitly pass a noop function, such as  _.noop or something similar.

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

- This is an example where we have a Promise-based implementation and we're converting it to allow both callbacks and Promises.
- Nodeify is used with the promise returned by `foo()`
- if callback is not a function, promise is returned as-is, otherwise callback will be called when promise is resolved or rejected

=====

### Dealing with concurrency

* `Promise.all([...])` executes all promises at once. How can we limit concurrency?

```js
const pAll = require('p-all')

const tasks = [
  task1,
  task2,
  task3,
  task4,
  task5
]

pAll(tasts, {concurrency: 2})
  .then(result => {
	  console.log(result)
  })
```

NOTES:
_[1 minutes]_

- One thing to note here is that we're actually passing an array of function tasks that when executed will return a Promise
- `Promise.all()` takes Promises
- `p-all` takes functions

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

- Similarly there is a `p-map` module for mapping. 
- `got` is a Promise based HTTP module

=====

### Use Native Promises

* Native Promises are awesome for compatibility and future-proofness
* Libraries like `bluebird` and `Q` add extras
* There is a whole ecosystem of tiny functional modules for native Promises
* https://github.com/wbinnssmith/awesome-promises
* In the real world Promise implementation is unlikely your performance bottleneck
* Native Promises will only get faster
* `async` / `await` returns a native Promise no matter what implementation you use

NOTES:
_[3 minutes]_

- Native Promises are awesome for compatibility and future-proofness
- Libraries like `bluebird` and `Q` mutate the constructor and add extras
- There is a whole ecosystem of tiny functional modules for native Promises
- https://github.com/wbinnssmith/awesome-promises
- A note on performance... In the real world Promise implementation is unlikely your performance bottleneck
- I can understand the argument that every tick counts, but in most practical scenarios there are more usually relevant performance considerations and issues.
- Native Promises in Node 7 and 8 are already pretty good (not as good as Bluebird) but they will only get faster
- Async / await returns a native Promise no matter what implementation you use, so even if you use a 3rd party library, you will still only get a Native Promise.

=====

### `async` / `await`

* The real exciting and _practical_ evolution of async Javascript
* An `async` function returns a Promise
* Use `await` on a Promise to wait for it be resolved or an exception thrown
* Just act on Promises under the hood

NOTES:
_[1 minutes]_

- Which brings us to Async /await. 
- Promises are just the necessary, not as elegant, stepping stone towards a better mechanism
- In my opinion this is real exciting and _practical_ evolution of async Javascript
* An `async` function returns a Promise
* We use `await` on a Promise to wait for it be resolved or an exception thrown
* Underneath they just act on Promises. Everything relevant to Promises still applies here.

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

* Executes just like it reads!

NOTES:
_[1 minutes]_

- Executes just like it reads!
- `await` can only be used within an `async` function
- Here when we use `await` it really just means stop this function and wait for whatever Promise is on the right to be resolved
- Event loop keeps on ticking servicing other functions, requests, whatever.

=====

### Benefits

* Even more concise and cleaner code over Promises
* Better error handling with normal try / catch
* Intuitive handling of conditional logic
* Better error stacks that points exactly to where the exception was thrown
* Easier debugging, step through code as if it was synchronous

NOTES:
_[3 minutes]_

- Even more concise and cleaner code over Promises
- Better error handling with normal try / catch. Async/await makes it finally possible to handle both synchronous and asynchronous errors with the same construct, good old try/catch.
- Intuitive handling of conditional logic. We can just use normal if / else constructs
- Better error stacks that points exactly to where the exception was thrown
- Easier debugging. With `async` and `await` you can step through `await` calls exactly as if they were normal synchronous calls.

=====

### Sequential async

```js
async function main () {
  const urls = ['github.com', 'standardjs.com', 'nodejs.org']
  for (let url of urls) {
    const msg = await get(url)
    console.log(msg)
  }
}

async function get(url) {
  const r = await got(url)
  return r.statusMessage
}

main()
// OK
// OK
// OK
```

NOTES:
_[1 minutes]_

- Doing equential async operations now becomes just a matter of using a for loop

=====

### Concurrent execution

```js
const got = require('got')

async function main () {
  const urls = ['github.com', 'standardjs.com', 'nodejs.org']
  const promises = urls.map(url => get(url))
  let results = await Promise.all(promises);
  console.log(results); // [ 'OK', 'OK', 'OK' ]
}

async function get(url) {
  const r = await got(url)
  return r.statusMessage
}

main() 
```

NOTES:
_[1 minutes]_

- Just use `Promise.all()`
- We could also do for of on the promises array and push the values into a results array but I think this code is better

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

NOTES:
_[1 minutes]_

- Some common pitfalls when using async functions
- It makes asynchronous code less obvious
- Not using `await`. If we don't use it we get just the Promise. Sometime this is useful. But it can also be a cause of bugs and issues.
- `await`-ing multiple values in a row is not very efficient. 
- We have to wait for first call to finish before doing the 2nd one. 
- Use `Promise.all()` unless needed or intended

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

- This is nicer, but it's _not great_. Still feels and looks a little awkward, and callbeck-y.

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

* With Promises and `async` / `await` we have a much more practical and elegant mechanism for asynchronous code.
* `async` / `await` is currently supported in Node.js Stable and most browsers
* Native Promise-based Node.js is coming eventually
* Start using Promises!

NOTES:
_[1 minutes]_

=====

### **Questions?**
