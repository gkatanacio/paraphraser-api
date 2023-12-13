# paraphraser-api

Backend for paraphrasing tool leveraging LLM APIs.

Implemented as an AWS Lambda application and utilizes Serverless Framework for deployment.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

## Endpoints

- POST /paraphrase
  - available providers: "chatgpt"
  - available tones: "formal", "amicable", "fun", "casual", "sympathetic", "persuasive"
  - sample request:
    ```
    POST /paraphrase
    {
        "provider": "chatgpt",
        "tone": "formal",
        "text": "I'm hungry. What's for dinner?"
    }
    ```
  - sample response:
    ```
    {
        "result": "I am currently experiencing hunger. May I inquire about the menu for this evening's meal?"
    }
    ```

## Usage

#### configure

```bash
$ make .env
```

- see generated `.env` file for configuration

#### tidy dependencies

```bash
$ make deps
```

#### run tests

```bash
$ make test
```

#### build serverless functions

```bash
$ make build
```

- this generates `bin` directory to be used in deployment

#### deploy serverless application

```bash
$ make deploy
```

### Helpers during development:

#### format all .go files in project (using go fmt)

```bash
$ make fmt
```

#### generate test mocks (to be used with [stretchr/testify](https://github.com/stretchr/testify)) for all interfaces in project

```bash
$ make mocks
```

- can be configured in `.mockery.yaml`
