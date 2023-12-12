import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { SoundeffectControllerComponent } from "./soundeffect-controller/soundeffect-controller.component";
import { UploadMp3Component } from "./upload-mp3/upload-mp3.component";

@Component({
    selector: 'app-root',
    standalone: true,
    templateUrl: './app.component.html',
    styleUrl: './app.component.css',
    imports: [CommonModule, RouterOutlet, SoundeffectControllerComponent, UploadMp3Component]
})
export class AppComponent {
  title = 'trein-controller-client';
}
