export const tokyoNight = {
  colors: {
    bg: {
      primary: '#1a1b26',
      secondary: '#16161e',
      tertiary: '#1f2335'
    },
    text: {
      primary: '#a9b1d6',
      secondary: '#787c99',
      accent: '#7aa2f7'
    },
    accent: {
      blue: '#7aa2f7',
      purple: '#9d7cd8',
      cyan: '#7dcfff',
      green: '#9ece6a',
      orange: '#ff9e64',
      red: '#f7768e'
    }
  },
  spacing: {
    sidebar: '250px',
    padding: {
      sm: '0.5rem',
      md: '1rem',
      lg: '2rem'
    }
  },
  transitions: {
    default: 'all 0.2s ease-in-out'
  }
};

export type Theme = typeof tokyoNight;
