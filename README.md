![TutorCare (1)](https://user-images.githubusercontent.com/59323055/150261046-70941ab4-8bed-46a0-a3ee-65d22ead7ddb.png) \
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
![Angular.js](https://img.shields.io/badge/angular.js-%23E23237.svg?style=for-the-badge&logo=angularjs&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
![MUI](https://img.shields.io/badge/MUI-%230081CB.svg?style=for-the-badge&logo=material-ui&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
# An application for on-demand childcare and tutoring services
## JIA-1307
[![GitHub issues](https://img.shields.io/github/issues/k-lombard/TutorCare)](https://github.com/k-lombard/TutorCare/issues)
[![GitHub forks](https://img.shields.io/github/forks/k-lombard/TutorCare)](https://github.com/k-lombard/TutorCare/network)
[![GitHub stars](https://img.shields.io/github/stars/k-lombard/TutorCare)](https://github.com/k-lombard/TutorCare/stargazers)
[![GitHub license](https://img.shields.io/github/license/k-lombard/TutorCare)](https://github.com/k-lombard/TutorCare/blob/main/LICENSE)
![GitHub top language](https://img.shields.io/github/languages/top/k-lombard/TutorCare)
[![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fk-lombard%2FTutorCare)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fk-lombard%2FTutorCare)


# Release Notes
## Version 1.0.0
### New Features
- Custom login and logout authentication
- Email Verification and GaTech Restriction
- Customizable public user profiles
- Map of available nearby caregivers
- Job Post: Create, Edit, Apply, Applications Received, Select
- Active Job: Start and end verification
- Ratings and Reviews
- Live Chat Messaging

### Bug Fixes
- Fixed explicit SQL queries with GORM refactoring to prevent SQL injection
- Put validation to prevent invalid information on the following forms: sign up, login, create job, and edit job
- Added authentication to backend of websockets to prevent hijacking

### Known Issues
- Webpack bundle needs to be optimized
- Search feature is not elasticsearch so it will become slow with higher number of posts
- Usability: Forms do not autofill information, requiring user to input the same info in multiple places
- Mobile sidebar is not closable on very small screen sizes

# Setup Guide for Development Server
### Required Pre-Requisite Installations
- Git (latest version) <br>
  https://git-scm.com/downloads
- Node version manager, nvm (macOS/linux) or nodist (Windows) (latest version) <br>
  https://docs.npmjs.com/downloading-and-installing-node-js-and-npm
- PostgreSQL Relational Database (latest version) <br>
  https://www.postgresql.org/download/
- Docker (latest version) <br>
  https://docs.docker.com/get-docker/
- Docker Desktop UI Application (latest version) <br>
  https://www.docker.com/products/docker-desktop/
- GoLang (latest version) <br>
  https://go.dev/doc/install

### Download Instructions
- Click on the green code button at the top right of this GitHub page
- Copy the HTTPS URL
- Open terminal or command prompt on your computer
- cd into the folder in which you want to download (clone) this project. The following is an example, it will be different on your device
```
  cd users/documents/Github/
```
- Run ``` git clone <HTTPS URL> ```  with the HTTPS URL of this project (you just copied)

### Install Required Node Modules
- cd into the 'client' folder. Example: ``` cd user/documents/github/tutorcare-core/client ```
- Run ``` npm install ```
- Run ``` npm install -g @angular/cli ``` 
- Run ``` npm install -D tailwindcss ```

### Starting Client Frontend
- Open new terminal or command prompt on your computer
- cd into the 'client' folder. Example: ``` cd user/documents/github/tutorcare-core/client ```
- **If on Windows run:**
```
  ng run ng serve --proxy-config proxy.config.json
```
- **If on MacOS/Linux run:**
```
  npm run start
```
- Keep this terminal/command prompt open

### Starting API Backend
- Open new terminal or command prompt on your computer
- cd into 'api' folder. Example: ``` cd user/documents/github/tutorcare-core/api ```
- **If on Windows run:**
```
  docker-compose up web
```
- followed by: 
```
  docker-compose up migrate
```
- **If on MacOS/Linux run:**
```
  sudo docker-compose up web
```
- followed by:
```
  sudo docker-compose up migrate
```
- Keep this terminal/command prompt open
