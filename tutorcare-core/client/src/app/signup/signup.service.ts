import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';

@Injectable()
export class SignupService {
  results:Object[];
  _output: any[] | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient, private toastr: ToastrService) {
    this.results = []
  }

  signup(firstName: string, lastName: string, email: string, password: string, user_category: string): Observable<Object[]> {
    // let url = `${environment.serverUrl}/api/signup/`;
    let url = `/api/signup/`;
    return new Observable((observer: any) => {
       this.http.post<any>(
            url,
            JSON.stringify({
              "first_name": firstName,
              "last_name": lastName,
              "email": email,
              "password": password,
              "user_category": user_category
            }),
            {headers: this.headers})
              .pipe(
                map((res: any) => res),
                catchError((err: HttpErrorResponse) => {
                  this.toastr.error("Error " + err.status + " " + err.error, "Error", {closeButton: true, timeOut: 5000, progressBar: true});
                  return throwError(err)
                })
              )
              .subscribe(
                (data: any) => {
                  this._output = data
                  observer.next(this._output);
                  observer.complete();
                },
                error => {return throwError(error)}
              );
    });
 }

 verifyCode(email: string, code: number): Observable<string> {
  // let url = `${environment.serverUrl}/api/signup/`;
  let url = `/api/signup/verify`;
  return new Observable((observer: any) => {
     this.http.post<any>(
          url,
          JSON.stringify({
            "email": email,
            "code": code
          }),
          {headers: this.headers}, )
            .pipe(
              map((res: any) => {
                console.log(res.status)}
              ),
              catchError((err: HttpErrorResponse) => {
                return throwError(err)
              })
            )
            .subscribe(
              (data: any) => {
                // console.log(data.status)
                this._output = data
                observer.next(this._output);
                observer.complete();
              },
              error => {return throwError(error)}
            );
    });
  }

  resendCode(email: string): Observable<string> {
    // let url = `${environment.serverUrl}/api/signup/`;
    console.log("resendEmailCode function")
    let url = `/api/signup/resendemail/${email}`;
    return new Observable((observer: any) => {
       this.http.post<any>(
            url,
            JSON.stringify({
              "email": email
            }),
            {headers: this.headers}, )
              .pipe(
                map((res: any) => {
                  console.log(res.status)}
                ),
                catchError((err: HttpErrorResponse) => {
                  return throwError(err)
                })
              )
      });
    }

}
