import { writable } from 'svelte/store';

export interface SettingsType {
  fontSize: string;
  color: string;
  fontFamily: string;
  backgroundColor?: string;
  backgroundColor2?: string;
  inputColor?: string;
  buttonColor?: string;
}

export const lightTheme: SettingsType = {
  fontSize: '30px',
  color: '#000',
  fontFamily: 'Arial',
  backgroundColor: '#FFF',
  backgroundColor2: '#F0F0F0',
  inputColor: '#E5E5E5',
  buttonColor: '#808080',
};

export const darkTheme: SettingsType = {
  fontSize: '30px',
  color: '#FFF',
  fontFamily: 'Arial',
  backgroundColor: '#292929',
  backgroundColor2: '#282828',
  inputColor: '#333',
  buttonColor: '#1abc9c'
};

export const highContrastTheme: SettingsType = {
fontSize: '30px',
color: '#FFF',
fontFamily: 'Arial',
backgroundColor: '#000',
backgroundColor2: '#202020',
inputColor: '#505050',
buttonColor: '#f1c40f'
};

export const darculaTheme: SettingsType = {
fontSize: '30px',
color: '#BABABA',
fontFamily: 'Arial',
backgroundColor: '#282a36',
backgroundColor2: '#44475a',
inputColor: '#6272a4',
buttonColor: '#8f3985'
};



export const settings = writable<SettingsType>(darkTheme);
