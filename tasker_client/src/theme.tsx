import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    mode: 'dark', // Dark mode to fit the background and contrast
    primary: {
      main: '#aa76f5', // Soft lavender for primary elements
    },
    secondary: {
      main: '#ff8a79', // Peachy pink for secondary actions
    },
    background: {
      default: 'linear-gradient(135deg, #010101 0%, #402b4c 100%)', // Gradient from dark to deep purple
      paper: '#010101', // Paper background, same as primary dark color
    },
    text: {
      primary: '#fbe3d7', // Primary text color (light cream)
      secondary: '#957d86', // Secondary text color (muted taupe)
    },
    action: {
      active: '#aa76f5', // Color for hover states, links, etc.
      hover: '#ff8a79', // Hover effects for buttons, cards
    },
    error: {
      main: '#ff8a79', // Error color using the peachy pink for contrast
    },
  },
  typography: {
    fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif', // Modern font
    h1: {
      fontSize: '2.5rem',
      color: '#fbe3d7', // Primary text color for headings
    },
    h2: {
      fontSize: '2rem',
      color: '#fbe3d7', // Primary text color for headings
    },
    h3: {
      fontSize: '1.75rem',
      color: '#fbe3d7', // Primary text color for headings
    },
    body1: {
      fontSize: '1rem',
      color: '#fbe3d7', // Primary body text color
    },
    body2: {
      fontSize: '0.875rem',
      color: '#957d86', // Secondary text color for body2
    },
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: '8px', // Rounded corners for a sleek look
          padding: '8px 16px', // Proper padding for buttons
          textTransform: 'none', // Keep text as is (no uppercase)
          fontWeight: 600,
        },
        containedPrimary: {
          backgroundColor: '#aa76f5', // Light lavender for primary button
          '&:hover': {
            backgroundColor: '#957d86', // Muted lavender on hover
          },
        },
        containedSecondary: {
          backgroundColor: '#ff8a79', // Peachy pink for secondary button
          '&:hover': {
            backgroundColor: '#aa76f5', // Light lavender on hover
          },
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          backgroundColor: '#402b4c', // Dark purple card background
          borderRadius: '16px', // Rounded corners for cards
          boxShadow: '0 4px 10px rgba(0, 0, 0, 0.5)', // Sleek shadow
        },
      },
    },
    MuiTypography: {
      styleOverrides: {
        root: {
          fontWeight: 400,
        },
      },
    },
    MuiAppBar: {
      styleOverrides: {
        root: {
          backgroundColor: '#010101', // Dark background for the AppBar
          boxShadow: 'none', // Remove AppBar shadow for a clean look
        },
      },
    },
    MuiTable: {
      styleOverrides: {
        root: {
          backgroundColor: '#010101', // Dark background for the entire table
          borderCollapse: 'collapse', // Collapse borders for a neat look
        },
      },
    },
    MuiTableCell: {
      styleOverrides: {
        root: {
          color: '#fbe3d7', // Light cream color for text
          borderBottom: '1px solid #957d86', // Muted taupe for row borders
          padding: '12px 16px', // Padding for cells
        },
        head: {
          color: '#fbe3d7', // Light cream for header text
          fontWeight: 'bold',
          backgroundColor: '#402b4c', // Dark purple background for header cells
        },
        body: {
          color: '#fbe3d7', // Body text in light cream
        },
      },
    },
    MuiTableHead: {
      styleOverrides: {
        root: {
          backgroundColor: '#402b4c', // Dark purple for table header
        },
      },
    },
    MuiTableRow: {
      styleOverrides: {
        root: {
          '&:hover': {
            backgroundColor: '#3e2a45', // Slightly lighter purple on hover for rows
          },
          '&:nth-of-type(odd)': {
            backgroundColor: '#2a1f2d', // Slightly lighter dark background for odd rows
          },
          '&:nth-of-type(even)': {
            backgroundColor: '#1e131e', // Even rows with a slightly darker background
          },
        },
      },
    },
    MuiTableContainer: {
      styleOverrides: {
        root: {
          backgroundColor: '#010101', // Container background
          borderRadius: '8px', // Rounded corners for the table container
          boxShadow: '0 4px 10px rgba(0, 0, 0, 0.5)', // Soft shadow for the table container
        },
      },
    },
    MuiList: {
      styleOverrides: {
        root: {
          padding: '0', // Remove default padding
          borderRadius: '8px', // Rounded corners for the list container
        },
      },
    },
    MuiListItem: {
      styleOverrides: {
        root: {
          padding: '10px 20px', // Padding for each list item
          '&:hover': {
            backgroundColor: '#2a1f2d', // Lighter purple on hover for items
          },
          '&.Mui-selected': {
            backgroundColor: '#402b4c', // Dark purple background when selected
          },
        },
      },
    },
    MuiListItemButton: {
      styleOverrides: {
        root: {
          padding: '10px 20px', // Padding for each list item
          '&:hover': {
            backgroundColor: '#2a1f2d', // Lighter purple on hover for items
          },
          '&.Mui-selected': {
            backgroundColor: '#402b4c', // Dark purple background when selected
          },
        },
      },
    },
    MuiListItemText: {
      styleOverrides: {
        primary: {
          color: '#fbe3d7', // Primary text color (light cream) for list item text
        },
        secondary: {
          color: '#957d86', // Muted taupe color for secondary text
        },
      },
    },
    MuiListItemIcon: {
      styleOverrides: {
        root: {
          color: '#aa76f5', // Lavender color for icons in the list
          minWidth: '36px', // Ensure icons aren't too large
        },
      },
    },
  },
});

export default theme;

