import { Injectable, Renderer2, RendererFactory2 } from '@angular/core';

import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';

import { ToastrService } from 'ngx-toastr';
import { OverlayContainer } from '@angular/cdk/overlay';


@Injectable()
export class ThemeService {
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  private renderer: Renderer2;
  constructor(private http: HttpClient, private toastr: ToastrService, private overlay: OverlayContainer, private rendererFactory: RendererFactory2) {
    this.renderer = rendererFactory.createRenderer(null, null);
  }

  setLight(){
    // Whenever the user explicitly chooses light mode
    const html = document.getElementsByTagName('html')[0];
    document.documentElement.classList.remove('dark')
    html.classList.remove('dark')
    this.overlay.getContainerElement().classList.remove('dark')
    localStorage.setItem('theme','light')

  }


setTheme() {
  if (localStorage.getItem('theme') === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
}

setDark() {
  const html = document.getElementsByTagName('html')[0];
  // Whenever the user explicitly chooses dark mode
  document.documentElement.classList.add('dark')
  html.classList.add('dark')
  localStorage.setItem('theme', 'dark')
  this.overlay.getContainerElement().classList.add('dark')
}

toggleTheme() {
  if (localStorage.getItem('theme') === 'dark') {
    const html = document.getElementsByTagName('html')[0];
    document.documentElement.classList.remove('dark')
    html.classList.remove('dark')
    this.overlay.getContainerElement().classList.remove('dark')
    localStorage.setItem('theme','light')
    this.overlay.getContainerElement().classList.add('dark')
    this.renderer.addClass(document.body, "dark");
  } else if (localStorage.getItem('theme') === 'light') {
    const html = document.getElementsByTagName('html')[0];
    document.documentElement.classList.add('dark')
    html.classList.add('dark')
    localStorage.setItem('theme', 'dark')
    this.overlay.getContainerElement().classList.add('dark')
    this.renderer.removeClass(document.body, "dark");
  } else {
    const html = document.getElementsByTagName('html')[0];
    document.documentElement.classList.add('dark')
    html.classList.add('dark')
    localStorage.setItem('theme', 'dark')
    this.overlay.getContainerElement().classList.add('dark')
    this.renderer.addClass(document.body, "dark");
  }
}

}
