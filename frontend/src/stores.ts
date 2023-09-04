import { writable } from 'svelte/store';

export interface SettingsType {
  fontSize: string;
  color: string;
  fontFamily: string;
  backgroundColor?: string;
  backgroundColor2?: string;
  inputColor?: string;
  buttonColor?: string;
  showInfoButton?: boolean;
}

export const lightTheme: SettingsType = {
  fontSize: '20px',
  color: '#000',
  fontFamily: 'Arial',
  backgroundColor: '#FFF',
  backgroundColor2: '#F0F0F0',
  inputColor: '#E5E5E5',
  buttonColor: '#808080',
  showInfoButton: true,
};

export const darkTheme: SettingsType = {
  fontSize: '20px',
  color: '#FFF',
  fontFamily: 'Arial',
  backgroundColor: '#292929',
  backgroundColor2: '#282828',
  inputColor: '#333',
  buttonColor: '#1abc9c',
  showInfoButton: true,
};

export const highContrastTheme: SettingsType = {
fontSize: '20px',
color: '#FFF',
fontFamily: 'Arial',
backgroundColor: '#000',
backgroundColor2: '#202020',
inputColor: '#505050',
buttonColor: '#f1c40f',
showInfoButton: true,
};

export const nordTheme: SettingsType = {
  fontSize: '20px',
  color: '#D8DEE9',
  fontFamily: 'Arial',
  backgroundColor: '#2E3440',
  backgroundColor2: '#3B4252',
  inputColor: '#4C566A',
  buttonColor: '#88C0D0',
  showInfoButton: true,
};

export const oceanTheme: SettingsType = {
  fontSize: '20px',
  color: '#333',
  fontFamily: 'Arial',
  backgroundColor: '#E0F2F1',
  backgroundColor2: '#B2DFDB',
  inputColor: '#80CBC4',
  buttonColor: '#00796B',
  showInfoButton: true,
};

export const sunnyTheme: SettingsType = {
  fontSize: '20px',
  color: '#333',
  fontFamily: 'Arial',
  backgroundColor: '#FFF9C4',
  backgroundColor2: '#FFECB3',
  inputColor: '#FFE082',
  buttonColor: '#FF5722',
  showInfoButton: true,
};

export const settings = writable<SettingsType>(darkTheme);
