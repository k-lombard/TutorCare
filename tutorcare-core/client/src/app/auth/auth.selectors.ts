import {createSelector} from '@ngrx/store';


export const selectAuthState = (state: any) => state.auth;


export const isLoggedIn = createSelector(
  selectAuthState,
  (auth: any) => auth.loggedIn
);

export const getCurrUser = createSelector(
    selectAuthState,
    (auth: any) => auth.user
);

export const isLoggedOut = createSelector(
  isLoggedIn,
  loggedIn => !loggedIn
);