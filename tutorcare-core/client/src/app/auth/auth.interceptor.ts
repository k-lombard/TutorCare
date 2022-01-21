import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
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
constructor(private authService: AuthService, private toastr: ToastrService, private store: Store<AppState>){}

    intercept(r: HttpRequest<any>, nextReq: HttpHandler): Observable<HttpEvent<any>> {
        try {
            r = r.clone({
                setHeaders: {
                'Content-Type' : 'application/json; charset=utf-8',
                'Accept'       : 'application/json',
                'Authorization': `Bearer ${this.authService.getAccessToken()}`,
                },
            });
            return nextReq.handle(r)
        } catch {
            this.store
            .pipe(
                select(isLoggedIn)
            ).subscribe(data => {
                console.log(data)
                this.isLoggedIn = data
            })
            if (this.isLoggedIn == false) {
                return nextReq.handle(r)
            } else {
                this.prevUser = this.authService.getCurrUser()
                this.accessToken = this.authService.getAccessToken()
                this._boolObservable = this.authService.isTokenValid(this.authService.getAccessToken())
                this._boolObservable.subscribe((data: any) => {
                    this.isValid = data
                    console.log(data)
                })  
                this.toastr.error("Error", "Error", {closeButton: true, timeOut: 5000, progressBar: true});
                this._refreshObservable = this.authService.refreshToken(this.authService.getRefreshToken());

                this._refreshObservable.subscribe((data: any) => {
                    this.output = data;
                    console.log(data)
                });
                if (this.output && this.output.access_token && this.output.refresh_token) {
                    this.new_access_token = this.output?.access_token || ""
                    this.new_refresh_token = this.output?.refresh_token || ""
                    this.prevUser.access_token = this.new_access_token
                    this.prevUser.refresh_token = this.new_refresh_token
                    this.store.dispatch(new Login({user: this.prevUser}));
                    r = r.clone({
                        setHeaders: {
                        'Content-Type' : 'application/json; charset=utf-8',
                        'Accept'       : 'application/json',
                        'Authorization': `Bearer ${this.new_access_token}`,
                        },
                    });
                }
                return nextReq.handle(r);
            }        
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
    }
}