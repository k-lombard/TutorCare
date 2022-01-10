import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { AuthService } from './auth.service';
import { Injectable } from '@angular/core';
import { Token } from '../models/token.model';
import { Store } from '@ngrx/store';
import { AppState } from '../reducers';
import { User } from '../models/user.model';
import { Login } from './auth.actions';


@Injectable()
export class AuthInterceptor implements HttpInterceptor {
new_access_token!: string
new_refresh_token!: string
_refreshObservable: Observable<Token> | undefined
output!: Token
prevUser!: User
constructor(private authService: AuthService, private store: Store<AppState>){}
    intercept(r: HttpRequest<any>, nextReq: HttpHandler): Observable<HttpEvent<any>> {
        this.prevUser = this.authService.getCurrUser()
        try {
            console.log(this.authService.getAccessToken())
            r = r.clone({
                setHeaders: {
                'Content-Type' : 'application/json; charset=utf-8',
                'Accept'       : 'application/json',
                'Authorization': `Bearer ${this.authService.getAccessToken()}`,
                },
            });

            return nextReq.handle(r);
        }
        catch {
            try {
                this._refreshObservable = this.authService.refreshToken(this.authService.getRefreshToken());

                this._refreshObservable.subscribe((data: any) => {
                    this.output = data;
                    console.log(data)
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
                });
                return nextReq.handle(r);
            } catch {
                return nextReq.handle(r);
            }
        } 
        
    }
}