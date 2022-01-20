import {ModuleWithProviders, NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {LoginComponent} from './login/login.component';
import {MatCardModule} from "@angular/material/card";
import { MatInputModule } from "@angular/material/input";
import {RouterModule} from "@angular/router";
import {ReactiveFormsModule, FormsModule} from "@angular/forms";
import {MatButtonModule} from "@angular/material/button";
import { StoreModule } from '@ngrx/store';
import {AuthService} from "./auth.service";
import * as fromAuth from './auth.reducer';
import {AuthGuard} from './auth.guard';
import { AuthEffects } from './auth.effects';
import { EffectsModule } from '@ngrx/effects';
import { NavbarComponent } from '../navbar/navbar.component';
import { SignupComponent } from '../signup/signup.component';
import { HomeComponent } from '../home/home.component';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { HttpModule } from '@angular/http';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatGridListModule } from '@angular/material/grid-list';
import { AccountComponent } from '../account/account.component';
import {MatSelectModule} from '@angular/material/select';
import { GoogleMapsModule } from '@angular/google-maps'
import { FindCareComponent } from '../find-care/find-care.component';
import { BarRatingModule } from "ngx-bar-rating";
import { EditProfileComponent } from '../account/edit-profile/edit-profile.component';
import { AuthInterceptor } from './auth.interceptor';
import { ToastrModule } from 'ngx-toastr';
import { FindJobsComponent } from '../find-jobs/find-jobs.component';

@NgModule({
    imports: [
        CommonModule,
        FormsModule,
        MatSelectModule,
        ReactiveFormsModule,
        MatInputModule,
        MatCardModule,
        HttpClientModule,
        MatButtonModule,
        ReactiveFormsModule,
        HttpModule,
        MatToolbarModule,
        MatSidenavModule,
        MatListModule,
        MatIconModule,
        BarRatingModule,
        MatGridListModule,
        RouterModule,
        GoogleMapsModule,
        ToastrModule,
        StoreModule.forFeature('auth', fromAuth.authReducer),
        EffectsModule.forFeature([AuthEffects]),

    ],
    declarations: [
        LoginComponent,
        NavbarComponent,
        SignupComponent,
        HomeComponent,
        AccountComponent,
        FindCareComponent,
        EditProfileComponent,
        FindJobsComponent
    ],
    exports: [LoginComponent]
})
export class AuthModule {
    static forRoot(): ModuleWithProviders<AuthModule> {
        return {
            ngModule: AuthModule,
            providers: [
              AuthService,
              AuthGuard,
              {
                provide : HTTP_INTERCEPTORS,
                useClass: AuthInterceptor,
                multi   : true,
              }
            ]
        }
    }
}
