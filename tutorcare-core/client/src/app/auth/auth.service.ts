
   
import {Injectable} from "@angular/core";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {User} from "../models/user.model";
import { environment } from 'src/environments/environment';
import { map } from 'rxjs/operators';
import { Store, select } from '@ngrx/store';
import { AppState } from '../reducers';
import { getCurrUser } from './auth.selectors';
import { Token } from '../models/token.model';




@Injectable()
export class AuthService {
    headers = new HttpHeaders({
        'Content-Type': 'application/json'
    });
    user!: User
    access_token!: string
    refresh_token!: string
    constructor(private http:HttpClient, private store: Store<AppState>) {
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
               
           }), {headers: this.headers})
    }

    refreshToken(refresh_token: string): Observable<Token> {
        // let url = `${environment.serverUrl}/api/login/`;
        let url = `/api/token/refresh`;
        return this.http.post<any>(url, JSON.stringify({
               "refresh_token": refresh_token
           }), {headers: this.headers})
    }

    isTokenValid(access_token: string): Observable<Object> {
        // let url = `${environment.serverUrl}/api/login/`;
        let url = `/api/token/valid`;
        let headers2 = new HttpHeaders({
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${access_token}`
        });
        return this.http.get(url, {headers: headers2})
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

    getAccessToken(): string {
        console.log(this.user.access_token)
        return this.user.access_token ? this.user.access_token : ""
    }
    getRefreshToken(): string {
        console.log(this.user.refresh_token)
        return this.user.refresh_token ? this.user.refresh_token : ""
    }
    getCurrUser(): User {
        return this.user
    }
}