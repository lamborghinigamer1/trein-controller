import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';

interface Sound {
  name: string;
}

@Component({
  selector: 'app-soundeffect-controller',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './soundeffect-controller.component.html',
  styleUrls: ['./soundeffect-controller.component.css']
})
export class SoundeffectControllerComponent implements OnInit {
  private urlgetsounds: string = 'http://localhost:8080/allsounds';
  private url: string = 'http://localhost:8080/playsound';
  public mp3s: string[] = [];

  ngOnInit() {
    this.getsounds();

    // Poll every 5 seconds to update the list of MP3s
    setInterval(() => {
      this.getsounds();
    }, 1000);
  }

  async getsounds() {
    try {
      const response = await fetch(this.urlgetsounds);
      const sounds: Sound[] = await response.json();

      // Extract the "name" property from each object in the array
      this.mp3s = sounds.map(sound => sound.name);
    } catch (error) {
      console.error('Error fetching sounds:', error);
    }
  }

  fetchurl(mp3file: string) {
    const url = this.url;
    const data = { key: mp3file };

    console.log(data);

    fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: new URLSearchParams(data).toString(),
    })
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.text();
      })
      .then(result => console.log(result))
      .catch(error => console.error('Error:', error));
  }
}
