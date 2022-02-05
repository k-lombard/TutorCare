import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { AuthService } from './auth.service';
import { Injectable, OnInit } from '@angular/core';
import { Token } from '../models/token.model';
import { select, Store } from '@ngrx/store';
import { AppState } from '../reducers';
import { User } from '../models/user.model';
import { Login } from './auth.actions';
import 'rxjs/add/operator/catch';
import { ToastrService } from 'ngx-toastr';
import { isLoggedIn } from './auth.selectors';
import { catchError, retry } from 'rxjs/operators';
import { Router } from '@angular/router';
import { throwError } from 'rxjs';
import 'rxjs/add/operator/do';


@Injectable()
export class AuthInterceptor implements HttpInterceptor {
new_access_token!: string
new_refresh_token!: string
_refreshObservable: Observable<Token> | undefined
_boolObservable: Observable<Object> | undefined
output!: Token
prevUser!: User
isValid!: boolean
accessToken!: string
isLoggedIn!: boolean;
constructor(private authService: AuthService, private toastr: ToastrService, private store: Store<AppState>, private router: Router){}


    handleAuthError(error: any) {
      console.log("error ", error);
      let msg = "";
      if (error !== undefined && typeof error === "string") {
        msg = error;
      } else if (error.error !== undefined && typeof error.error === "string") {
        msg = error.error;
      } else if (
        error.error.error !== undefined &&
        typeof error.error.error === "string"
      ) {
        msg = error.error.error;
      } else {
        msg = error.error.error.errors
          ? error.error.error.errors[0].errorMessage
          : "Something went wrong";
      }
      this.toastr.error(msg, "", {
        timeOut: 3000,
        positionClass: "toast-bottom-center",
      });
    }


    // intercept(r: HttpRequest<any>, nextReq: HttpHandler): Observable<HttpEvent<any>> {
    //   try {
    //     r = r.clone({
    //         setHeaders: {
    //         'Content-Type' : 'application/json; charset=utf-8',
    //         'Accept'       : 'application/json',
    //         'Authorization': `Bearer ${this.authService.getAccessToken()}`,
    //         },
    //     });
    //     // return nextReq.handle(r)
    //     return nextReq.handle(r).pipe(retry(1), catchError((error: HttpErrorResponse) => {
    //       if (error.error instanceof HttpErrorResponse) {
    //         if(error.error instanceof ErrorEvent) {
    //           console.log("Error Event")
    //         } else {
    //           console.log(`error status : ${error.status} ${JSON.stringify(error.error)}`);
    //           switch (error.status) {
    //             case 401:
    //               this.router.navigateByUrl("/login");
    //               break;
    //             case 403:
    //               this.router.navigateByUrl("/unauthorized");
    //               break;
    //             case 0:
    //             case 400:
    //             case 405:
    //             case 406:
    //             case 409:
    //             case 500:
    //               this.handleAuthError(error);
    //               break;
    //             }
    //           }
    //     } else {
    //       console.error("something else haappened");
    //     }
    //   return throwError(error) }))}
    //     } catch {
    //         this.store
    //         .pipe(
    //             select(isLoggedIn)
    //         ).subscribe(data => {
    //             console.log(data)
    //             this.isLoggedIn = data
    //         })
    //         if (this.isLoggedIn == false) {
    //             return nextReq.handle(r)
    //         } else {
    //             // this.prevUser = this.authService.getCurrUser()
    //             // this.accessToken = this.authService.getAccessToken()
    //             // this._boolObservable = this.authService.isTokenValid(this.authService.getAccessToken())
    //             // this._boolObservable.subscribe((data: any) => {
    //             //     this.isValid = data
    //             //     console.log(data)
    //             // })
    //             this.toastr.error("Error", "Error", {closeButton: true, timeOut: 5000, progressBar: true});
    //             this._refreshObservable = this.authService.refreshToken(this.authService.getRefreshToken());

    //             this._refreshObservable.subscribe((data: any) => {
    //                 this.output = data;
    //                 console.log(data)
    //             });
    //             if (this.output && this.output.access_token && this.output.refresh_token) {
    //                 this.new_access_token = this.output?.access_token || ""
    //                 this.new_refresh_token = this.output?.refresh_token || ""
    //                 this.prevUser.access_token = this.new_access_token
    //                 this.prevUser.refresh_token = this.new_refresh_token
    //                 this.store.dispatch(new Login({user: this.prevUser}));
    //                 r = r.clone({
    //                     setHeaders: {
    //                     'Content-Type' : 'application/json; charset=utf-8',
    //                     'Accept'       : 'application/json',
    //                     'Authorization': `Bearer ${this.new_access_token}`,
    //                     },
    //                 });
    //             }
    //             return nextReq.handle(r);
    //         }
    //     }






        intercept(r: HttpRequest<any>, nextReq: HttpHandler): Observable<HttpEvent<any>> {
          if (this.authService.isLoggedInFunc()) {
            r = r.clone({
                setHeaders: {
                'Content-Type' : 'application/json; charset=utf-8',
                'Accept'       : 'application/json',
                'Authorization': `Bearer ${this.authService.getAccessToken()}`,
                },
            });

            return nextReq.handle(r).do((event: HttpEvent<any>) => {
              if (event instanceof HttpResponse) {
                // do stuff with response if you want
                console.log("Authorized")
              }
            }, (err: any) => {
              if (err instanceof HttpErrorResponse) {
                if (err.status === 401 && this.authService.isLoggedInFunc()) {
                  this._refreshObservable = this.authService.refreshToken(this.authService.getRefreshToken());
                  this._refreshObservable.subscribe((data: any) => {
                      this.output = data;
                      console.log(data)
                    if (this.output && this.output.access_token && this.output.refresh_token) {
                        this.new_access_token = this.output?.access_token || ""
                        this.new_refresh_token = this.output?.refresh_token || ""
                        this.prevUser.access_token = this.new_access_token
                        this.prevUser.refresh_token = this.new_refresh_token
                        this.store.dispatch(new Login({user: this.prevUser}));
                        this.toastr.success("Successfully refreshed access token.", "Success", {closeButton: true, timeOut: 5000, progressBar: true});
                    }
                  });
                  return
                } else {
                  this.router.navigate(['/login'])
                  this.toastr.error("Refresh token expired. Could not refresh access token.", "Error", {closeButton: true, timeOut: 5000, progressBar: true});
                }
              }
          })
        } else {
          return nextReq.handle(r)
        }
      }
          // return nextReq.handle(r)
          // return nextReq.handle(r).pipe(retry(1), catchError((error: HttpErrorResponse) => {
          //   if (error.error instanceof HttpErrorResponse) {
          //     if(error.error instanceof ErrorEvent) {
          //       console.log("Error Event")
          //     } else {
          //       console.log(`error status : ${error.status} ${JSON.stringify(error.error)}`);
          //       switch (error.status) {
          //         case 401:
          //           this._refreshObservable = this.authService.refreshToken(this.authService.getRefreshToken());
          //           this._refreshObservable.subscribe((data: any) => {
          //               this.output = data;
          //               console.log(data)
          //           });
          //           if (this.output && this.output.access_token && this.output.refresh_token) {
          //               this.new_access_token = this.output?.access_token || ""
          //               this.new_refresh_token = this.output?.refresh_token || ""
          //               this.prevUser.access_token = this.new_access_token
          //               this.prevUser.refresh_token = this.new_refresh_token
          //               this.store.dispatch(new Login({user: this.prevUser}));
          //               r = r.clone({
          //                   setHeaders: {
          //                   'Content-Type' : 'application/json; charset=utf-8',
          //                   'Accept'       : 'application/json',
          //                   'Authorization': `Bearer ${this.new_access_token}`,
          //                   },
          //               });
          //           }
          //           return nextReq.handle(r).pipe(retry(1), catchError((error: HttpErrorResponse) => {
          //             if (error.error instanceof HttpErrorResponse) {
          //               if(error.error instanceof ErrorEvent) {
          //                 console.log("Error Event")
          //               } else {
          //                 console.log(`error status : ${error.status} ${JSON.stringify(error.error)}`);
          //                 switch (error.status) {
          //                   case 401:
          //                     this.router.navigateByUrl("/login");
          //                     break;
          //                   case 403:
          //                   case 0:
          //                   case 400:
          //                   case 405:
          //                   case 406:
          //                   case 409:
          //                   case 500:
          //                     this.handleAuthError(error);
          //                     break;
          //                   }
          //                 }
          //           } else {
          //             console.error("something else haappened");
          //           }
          //           return throwError(error) }))
          //         case 403:
          //         case 0:
          //         case 400:
          //         case 405:
          //         case 406:
          //         case 409:
          //         case 500:
          //           this.handleAuthError(error);
          //           break;
          //         }
          //       }
          // } else {
          //   console.error("something else haappened");
          // }
          // return throwError(error) }))}

  }

















        // .catch((err: any) => {
        //     if (err.status == 401) {
        //     this.toastr.error(err.error.errors.name, "Error", {closeButton: true, timeOut: 5000, progressBar: true});
        //         this._refreshObservable = this.authService.refreshToken(this.authService.getRefreshToken());

        //         this._refreshObservable.subscribe((data: any) => {
        //             this.output = data;
        //             console.log(data)
        //         });
        //         if (this.output && this.output.access_token && this.output.refresh_token) {
        //             this.new_access_token = this.output?.access_token || ""
        //             this.new_refresh_token = this.output?.refresh_token || ""
        //             this.prevUser.access_token = this.new_access_token
        //             this.prevUser.refresh_token = this.new_refresh_token
        //             this.store.dispatch(new Login({user: this.prevUser}));
        //             r = r.clone({
        //                 setHeaders: {
        //                 'Content-Type' : 'application/json; charset=utf-8',
        //                 'Accept'       : 'application/json',
        //                 'Authorization': `Bearer ${this.new_access_token}`,
        //                 },
        //             });
        //         }
        //         return nextReq.handle(r);
        //     }
        //     return nextReq.handle(r)
        // })
