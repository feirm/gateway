<h1 align="center">Feirm Gateway</h1>
<p align="center">⛩️ API Gateway for Feirm Microservices.</p>

<p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/feirm/gateway?style=for-the-badge" alt="Go Version" />
    <img src="https://goreportcard.com/badge/github.com/feirm/gateway?style=for-the-badge" alt="Go Report Card"/>
</p>

The Feirm Gateway essentially acts as a router to our microservices. It is configured through a JSON configuration file and also enforces rate limiting to prevent services from being abused.

You can find this gateway being used in production for Feirm at [https://api.feirm.com](https://api.feirm.com)