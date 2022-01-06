
   
import {Injectable} from "@angular/core";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {User} from "../models/user.model";
import { environment } from 'src/environments/environment';
import { map } from 'rxjs/operators';




@Injectable()
export class AuthService {
    headers = new HttpHeaders({
        'Content-Type': 'application/json'
    });
    constructor(private http:HttpClient) {

    }

    // login(email:string, password:string): Observable<User> {
    //     return this.http.post<User>('/api/login', {email,password});
    // }

    login(email: string, password: string): Observable<User> {
        // let url = `${environment.serverUrl}/api/login/`;
        let url = `/api/login/`;
        return this.http.post<any>(url, JSON.stringify({
               "first_name": "",
               "last_name": "",
               "email": email,
               "password": password,
               
           }), {headers: this.headers})
    }

    // getPosition(): Observable<any> {
    //     return new Observable((observer: any) => {
    //       window.navigator.geolocation.getCurrentPosition(position => {
    //         observer.next(position);
    //         observer.complete();
    //       },
    //         error => observer.error(error));
    //     });
    // }
}