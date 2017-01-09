<?php

require_once __DIR__.'/vendor/autoload.php';

use Pheanstalk\Pheanstalk;
use Symfony\Component\Console\Application;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Console\Input\InputOption;

$faker = \Faker\Factory::create();

$app = new Application('fixtures', '1.0.0');
$app
    ->register('fixtures')
    ->addOption('host', null, InputOption::VALUE_OPTIONAL, 'The beanstalkd host', '0.0.0.0')
    ->addOption('num', null, InputOption::VALUE_OPTIONAL, 'Number of fixtures', 10)
    ->setCode(function(InputInterface $input, OutputInterface $output) use ($faker) {
        $host = $input->getOption('host');
        $n = $input->getOption('num');

        $pheanstalk = new Pheanstalk($host);

        for ($i = 0; $i < $n; $i++) {
            $fixture = [
                'userId' => $faker->randomNumber(7),
                'preferences' => [
                    $faker->unique($reset = true)->randomDigitNotNull,
                    $faker->unique()->randomDigit,
                    $faker->unique()->randomDigit,
                ],
            ];

            $pheanstalk->put(json_encode($fixture));
            $output->writeln('Queued fixture #'.($i+1));
        }
    })
    ->getApplication()
    ->setDefaultCommand('fixtures', true)
    ->run()
;
