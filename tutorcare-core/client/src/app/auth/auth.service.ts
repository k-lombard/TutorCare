

import {Injectable} from "@angular/core";
import {HttpClient, HttpErrorResponse, HttpHeaders} from "@angular/common/http";
import {Observable, throwError} from "rxjs";
import {User} from "../models/user.model";
import { environment } from 'src/environments/environment';
import { catchError, map, tap } from 'rxjs/operators';
import { Store, select } from '@ngrx/store';
import { AppState } from '../reducers';
import { getCurrUser, isLoggedIn } from './auth.selectors';
import { Token } from '../models/token.model';
import { GeolocationPosition } from "../models/geolocationposition.model";
import { ToastrService } from "ngx-toastr";
import { Router } from "@angular/router";
import { Login, Logout } from "./auth.actions";



@Injectable()
export class AuthService {
    headers = new HttpHeaders({
        'Content-Type': 'application/json'
    });
    user!: User
    _isValid: boolean | undefined
    access_token!: string
    refresh_token!: string
    isLoggedIn: boolean
    output: Token
    new_access_token: string
    new_refresh_token: string
    prevUser: User
    constructor(private http:HttpClient, private store: Store<AppState>, private toastr: ToastrService, private router: Router) {
      this.store
      .pipe(
          select(getCurrUser)
      ).subscribe(data =>  {
          this.prevUser = data
    })
      this.store
      .pipe(
        select(isLoggedIn)
      ).subscribe(data => {
        this.isLoggedIn = data
      })
        this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            try {
                this.access_token = this.user.access_token ? this.user.access_token : ""
                this.refresh_token = this.user.refresh_token ? this.user.refresh_token : ""
            } catch {
                this.access_token = ""
                this.refresh_token = ""
            }
        })
    }

    // login(email:string, password:string): Observable<User> {
    //     return this.http.post<User>('/api/login', {email,password});
    // }

    login(email: string, password: string): Observable<User> {
        // let url = `${environment.serverUrl}/api/login/`;
        let url = `/api/login/`;
        return this.http.post<any>(url, JSON.stringify({
               "email": email,
               "password": password,

           }), {headers: this.headers}).pipe(
            map((res: any) => res),
            catchError((err: HttpErrorResponse) => {
              this.toastr.error("Error logging in.", "Error", {closeButton: true, timeOut: 5000, progressBar: true});
              return throwError(err)
            })
          )
    }

    logout(): Observable<any> {
      // let url = `${environment.serverUrl}/api/login/`;
      let url = `/api/logout/`;
      return this.http.post<any>(url, JSON.stringify({}), {headers: this.headers}).pipe(
          map((res: any) => res),
          catchError((err: HttpErrorResponse) => {
            this.toastr.error("Error logging out.", "Error", {closeButton: true, timeOut: 5000, progressBar: true});
            return throwError(err)
          })
        )
  }

    refreshToken(refresh_token: string): Observable<Token> {
        // let url = `${environment.serverUrl}/api/login/`;
        console.log(refresh_token)
        let url = `/api/token/refresh`;
        return new Observable((observer: any) => {
          this.http.post<any>(url, JSON.stringify({
               "refresh_token": refresh_token
           }), {headers: this.headers})
          .pipe(map((res: Token) => res),
          catchError((err: HttpErrorResponse) => {
            this.toastr.error("Error refreshing access token. Please log back in.", "Error", {closeButton: true, timeOut: 5000, progressBar: true});
            this.logout()
            .pipe(
              tap(user => {
                this.store.dispatch(new Logout());
              })
            )
            this.router.navigate['/login']
            return throwError(err)
          })
          )
          .subscribe((data: Token) => {
            console.log(this.prevUser.access_token)
            console.log(this.prevUser.refresh_token)
            this.output = data
            console.log(this.output)
            if (this.output && this.output.access_token && this.output.refresh_token) {
              let newUser = {...this.prevUser}
              this.new_access_token = this.output?.access_token || ""
              this.new_refresh_token = this.output?.refresh_token || ""
              newUser.access_token = this.new_access_token
              newUser.refresh_token = this.new_refresh_token
              console.log(newUser)
              this.store.dispatch(new Login({user: newUser}));
              this.toastr.success("Successfully refreshed access token.", "Success", {closeButton: true, timeOut: 5000, progressBar: true});
            }
            observer.next(data);
            observer.complete();
          })
        })
    }

    isLoggedInFunc() {
      return this.isLoggedIn
    }

    isTokenValid(access_token: string): Observable<Object> {
        // let url = `${environment.serverUrl}/api/login/`;
        let url = `/api/token/valid`;
        let headers2 = new HttpHeaders({
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${access_token}`
        });
        return new Observable((observer: any) => {
            this.http.get(url, {headers: headers2})
                .pipe(map((res: any) => res.json()))
                .subscribe((data: any) => {
                    this._isValid = data
                    observer.next(this._isValid);
                    observer.complete();
                });
         });
    }

    getPosition(): Observable<any> {
        return new Observable((observer: any) => {
          window.navigator.geolocation.getCurrentPosition(position => {
            observer.next(position);
            observer.complete();
          },
            error => observer.error(error));
        });
    }

    createPosition(user_id: string, accuracy: number, latitude: number, longitude: number): Observable<GeolocationPosition> {
      let url = `/api/geolocationpositions/`;
        return this.http.post<any>(url, JSON.stringify({
               "user_id": user_id,
               "accuracy": accuracy,
               "latitude": latitude,
               "longitude": longitude
           }), {headers: this.headers})
    }

    getAccessToken(): string {
        return this.user.access_token ? this.user.access_token : ""
    }
    getRefreshToken(): string {
        return this.user.refresh_token ? this.user.refresh_token : ""
    }
    getCurrUser(): User {
        return this.user
    }
}
