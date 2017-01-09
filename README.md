# recommender

A simple worker pool POC for [beanstalkd](http://kr.github.io/beanstalkd/) written in Go.

## Quickstart

Build all Docker containers

```bash
bin/build/beanstalkd
bin/build/recommender
```

Start containers

```bash
bin/app/start
```

You should now see some spawned workers ready to process incoming beanstalkd jobs:

```bash
Starting worker #1
Starting worker #2
Starting worker #3
Starting worker #4
Starting worker #5
Starting worker #6
Starting worker #7
Starting worker #8
Starting worker #9
Starting worker #10
```

### Testing (Creating sample jobs)

There's a dead simple PHP client included to put jobs on the beanstalkd queue.

Open another terminal window, and install Composer dependencies

```bash
cd test/fixtures
bin/composer install
```

Create some beanstalkd jobs:

```bash
bin/app/test [--num=100]
```

The option `--num=<n>` will create `<n>` jobs (defaults to 10)

The workers will now start processing the jobs:

```bash
Add job #1 (Body: {"userId":1834515,"preferences":[9,5,1]})
Processing job with ID #1
Add job #2 (Body: {"userId":6717618,"preferences":[3,0,2]})
Processing job with ID #2
Add job #3 (Body: {"userId":8048923,"preferences":[5,2,0]})
Processing job with ID #3
Add job #4 (Body: {"userId":79531,"preferences":[2,7,1]})
Processing job with ID #4
Add job #5 (Body: {"userId":6314766,"preferences":[1,0,6]})
Processing job with ID #5
Add job #6 (Body: {"userId":3145768,"preferences":[4,4,7]})
Processing job with ID #6
Add job #7 (Body: {"userId":5933815,"preferences":[4,9,3]})
Processing job with ID #7
Add job #8 (Body: {"userId":2917025,"preferences":[3,0,4]})
Processing job with ID #8
Add job #9 (Body: {"userId":667122,"preferences":[6,8,4]})
Processing job with ID #9
Add job #10 (Body: {"userId":4716235,"preferences":[6,4,8]})
Processing job with ID #10
Worker: #9: Finished processing job #10
Worker: #10: Finished processing job #1
Worker: #1: Finished processing job #2
Worker: #2: Finished processing job #3
Worker: #3: Finished processing job #4
Worker: #4: Finished processing job #5
Worker: #5: Finished processing job #6
Worker: #6: Finished processing job #7
Worker: #7: Finished processing job #8
Worker: #8: Finished processing job #9
```