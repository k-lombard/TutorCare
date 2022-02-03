import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';

@Injectable()
export class VerifyService {
  results:Object[];
  _output: any[] | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient, private toastr: ToastrService) {
    this.results = []
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
    console.log("resendEmailCode function")
    console.log(email)
    let url = `/api/signup/resendemail/${email}`;
    return new Observable((observer: any) => {
       this.http.get(
            url).pipe(
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

}
