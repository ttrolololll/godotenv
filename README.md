# GoDotenv

A stupidly simple (almost naive) implementation of <em>dotenv</em>.

## Motivation

I just wanted a quick way to load values from <em>.env</em> for project prototypes and POCs.

## Usage

<code>var env = dotenv.GetInstance()

var PRIVATETOKEN  = env.Env("GITLAB_TOKEN")</code>

## Disclaimer

Not for use in production.