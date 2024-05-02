How to run stack locally ?
Run following command in root of this project to spin up stack locally :
docker compose up

Workflows for all three technologies are defined in .github/workflows.

Each of the three workflows are separate and dedicated pipelines for each technology.

A workflow runs only if there was change in that technology's folder or workflow file itself was modified.
For example, Go workflow ( CI pipeline ) will only run if either any files in "go_component" were modified or .github/workflows/Go-CI.yaml was itself modified.

Why deploy stage is not included in any workflow ?
I have not included deploy stage because of lack of source code. If source code for some real life application was provided for each of the technologies, then it would have easier to write deploy stage. But since source code was not provided, I would first have to develop source code that would have taken significant time. ( Especially in case of Wordpress because I have not done any development in Wordpress, though I have worked on it from GUI and know its fundamentals. )

Please look at workflow files for any other details.
