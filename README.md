![TutorCare (1)](https://user-images.githubusercontent.com/59323055/150261046-70941ab4-8bed-46a0-a3ee-65d22ead7ddb.png)
## An application for on-demand childcare and tutoring services

[![GitHub issues](https://img.shields.io/github/issues/k-lombard/TutorCare)](https://github.com/k-lombard/TutorCare/issues)
[![GitHub forks](https://img.shields.io/github/forks/k-lombard/TutorCare)](https://github.com/k-lombard/TutorCare/network)
[![GitHub stars](https://img.shields.io/github/stars/k-lombard/TutorCare)](https://github.com/k-lombard/TutorCare/stargazers)
[![GitHub license](https://img.shields.io/github/license/k-lombard/TutorCare)](https://github.com/k-lombard/TutorCare/blob/main/LICENSE)
[![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fk-lombard%2FTutorCare)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fk-lombard%2FTutorCare)


### Starting Client Frontend
- cd into client folder
- **If on Windows run:**
```
  ng serve --proxy-config proxy.config.json
```
- **If on macOS/linux run:**
```
  npm run start
```

### Starting API Backend
- cd into api folder
- **If on Windows run:**
```
  docker-compose up web
```
- followed by: 
```
  docker-compose up migrate
```
- **If on macOS/linux run:**
```
  sudo docker-compose up web
```
- followed by:
```
  sudo docker-compose up migrate
```
