import test from 'ava';
import request from 'supertest-as-promised';

function getRandomInt(min, max) {
  return Math.floor(Math.random() * (max - min)) + min;
}

test('POST /users', async t => {
  t.plan(2);

  const n = getRandomInt(1, 1000);
  const data = {
    email: `dbojan+test+${n}@gmail.com.com`,
    firstName: `Test ${n}`,
    lastName: `User ${n}`
  };

  const res = await request('http://localhost:3000')
    .post('/users')
    .send(data);

  t.is(res.status, 201);
  t.deepEqual(res.body, data);
});
