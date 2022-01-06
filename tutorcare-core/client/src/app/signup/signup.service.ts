import { Injectable } from '@angular/core';
import {Http, RequestMethod, RequestOptions} from '@angular/http';
import {environment} from '../../environments/environment';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable() 
export class SignupService {
  results:Object[];
  _output: any[] | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
});
  constructor(private http: HttpClient) { 
    this.results = []
  }

  signup(firstName: string, lastName: string, email: string, password: string): Observable<Object[]> {
    let url = `${environment.serverUrl}/api/signup/`;
    return new Observable((observer: any) => {
       this.http.post<any>(url, JSON.stringify({
           "first_name": firstName,
           "last_name": lastName,
           "email": email,
           "password": password,
           
       }), {headers: this.headers})
           .pipe(map((res: any) => res.json()))
           .subscribe((data: any) => {
              this._output = data
 
              observer.next(this._output);
              observer.complete();
 
 
           });
    });
 }

}