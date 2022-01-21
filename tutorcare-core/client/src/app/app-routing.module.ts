import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './auth/login/login.component';
import { SignupComponent } from './signup/signup.component';
import { AuthModule } from './auth/auth.module';
import { AccountComponent } from './account/account.component';
import { FindCareComponent } from './find-care/find-care.component';
import { EditProfileComponent } from './account/edit-profile/edit-profile.component';
import { AuthGuard } from './auth/auth.guard';
import { CreateJobDialog, FindJobsComponent } from './find-jobs/find-jobs.component';

const routes: Routes = [
  { path: 'home', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'find-care', component: FindCareComponent },
  { path: 'find-jobs', component: FindJobsComponent },
  { path: 'about-us', component: HomeComponent },
  { path: 'account', component: AccountComponent, canActivate: [AuthGuard]},
  { path: 'account/edit-profile', component: EditProfileComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/create', component: CreateJobDialog, canActivate: [AuthGuard]},
  { path: '',   redirectTo: '/home', pathMatch: 'full' },
  { path: '**', redirectTo: '/home', pathMatch: 'full' },
  { path: '',   component: HomeComponent }

];

@NgModule({
  imports: [RouterModule.forRoot(routes), AuthModule],
  exports: [RouterModule]
})
export class AppRoutingModule { }
