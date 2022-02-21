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
## Version 0.2.0
### New Features
- Adds live chat messaging system for user to communicate to each other.
- Adds search and filter functionality to narrow the jobs shown on the find-jobs page.
- Adds applying to a job post functionality where a caregiver can apply to a job post.
- Adds editing and deleting job post functionality to allow change to be made to a job post.
- Adds functionality for a careseeker to view a list of applicants to their job post, see their profiles, and add them to a chatroom for messages.
- Adds hiring functionality for careseekers to hire someone from the list of applicant.
- Adds working find-care functionality for caregseekers to find nearby caregivers.
- Adds logout capability for a user.

### Bug Fixes
- Fixed bug where messages link did not show in mobile side navigation
- Fixed explicit SQL queries with GORM refactoring to prevent SQL injection
- Put validation to prevent invalid information on the following forms: sign up, login, create job, and edit job

### Known Issues
- Webpack bundle needs to be optimized
- Search feature is not elastic so it will be too slow with a larger number of posts
- Find-Care, Active Jobs, Applications Revcieved, Applications Sent pages are not mobile responsive
- Some forms do not autofill information, requiring user to input the same info in multiple places
- Mobile sidebar is not closable on very small screen sizes. Filter button clips in front of the nav on find-jobs page.

## Version 0.1.0
### New Features
- Adds working signup/login functionality and stores authentication tokens in session storage
- Adds working find-care page functionality where caregiver users are displayed in a user's local area
- Adds working find-jobs page functionality where job posts are displayed in an easy to see, intuitive manner
- Adds account page and edit profile functionality, to change a user's email, bio, experience, and user-type. 

### Known Issues
- Backend uses explicit SQL queries so it is vulnerable to SQL injection
- Chatroom doesn't scroll to bottom on page load
- Find-Care page's grid layout is not mobile responsive and does not work past a certain number of caregiver users


## Setup Guide
### Required Pre-Setup Installations
- Node version manager such as nvm for macOS/linux or nodist for windows (latest version)
- Install the latest node version and activate it with (nvm use "version")
- PostgreSQL (latest version)
- Docker (latest version)
- Go (latest version)

### Required Node_Modules Installations
- cd into the client folder
- Run ``` npm i ```
- Run ``` npm install -g @angular/cli ``` 
- Run ``` npm install -D tailwindcss ```

### Starting Client Frontend
- cd into client folder
- **If on Windows run:**
```
  ng run start-windows
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
