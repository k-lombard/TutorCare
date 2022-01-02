import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {environment} from '../environments/environment';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';


@Injectable() 
export class UsersService {
  results:Object[];
  _users: any[] | undefined;
  constructor(private http: Http) { 
    this.results = []
  }

  // users(): Observable<Object[]>{
  //   let ob = new Observable((observer: Object[]) => {
  //     let apiURL = `${environment.serverUrl}/api/users/`;
  //     this.http.get(apiURL)
  //       .toPromise()
  //       .then(
  //         (res: any) => { // Success
  //         this.results = res.json().users;
  //         console.log(this.results)
  //         observer.next(res.json().users);
  //         },
  //         (msg: any) => { // Error
  //         observer.error(msg);
  //         }
  //       );
  //   });
  //   return ob;
  // }

  getUsers(): Observable<Object[]> {
    let url = `${environment.serverUrl}/api/users/`;
 
    return new Observable((observer: any) => {
       this.http.get(url)
           .pipe(map((res: any) => res.json()))
           .subscribe((data: any) => {
              this._users = data
 
              observer.next(this._users);
              observer.complete();
 
 
           });
    });
 }

}