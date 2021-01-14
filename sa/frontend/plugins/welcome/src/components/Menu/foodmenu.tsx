import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Link, Container, makeStyles, Button, TextField,
  FormControl, InputLabel, Select, MenuItem,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
//import SaveIcon from '@material-ui/icons/Save';
//import { EntMenu } from '../../api/models/EntMenu';
import { EntMenugroup } from '../../api/models/EntMenugroup';
import { EntMenutype } from '../../api/models/EntMenutype';
import { EntUser } from '../../api/models/EntUser';

import { Content, Header, Page, pageTheme, } from '@backstage/core';
import { Alert } from '@material-ui/lab';

//-----------------------------------------------------------------------------

const lightColor = 'rgba(255, 255, 255, 0.7)';

const HeaderCustom = {
  minHeight: '100px',
};

const useStyles = makeStyles(theme => ({
  root: {
    '& > *': {
      marginRight: theme.spacing(3),
    },
  },
  form: {
    '& > *': {
      width: '100%',
      margin: theme.spacing(1),
    },
  },
  formControl: {
    margin: theme.spacing(1),
    Width: '100%',
  },
  textField: {
    margin: theme.spacing(1),
    width: '100%',
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  button: {
    borderColor: lightColor,
  },
}));

//-------------------------------------------------------------------------------
export default function CreateMenu() {

  const classes = useStyles();
  const api = new DefaultApi();

  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [menutypes, setMenutypes] = useState<EntMenutype[]>([]);
  const [menugroups, setMenugroups] = useState<EntMenugroup[]>([]);

  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
    };
    getUsers();

    const getMenugroup = async () => {
      const res = await api.listMenugroup({ limit: 10, offset: 0 });
      setLoading(false);
      setMenugroups(res);
    };
    getMenugroup();

    const getMenutype = async () => {
      const res = await api.listMenutype({ limit: 10, offset: 0 });
      setLoading(false);
      setMenutypes(res);
    };
    getMenutype();

  }, [loading]);


  const UserIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetOwnerID(event.target.value as number);
  };

  const MenutypeIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetTypeID(event.target.value as number);
  };

  const MenugroupIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetGroupID(event.target.value as number);
  };

  const Menuname_handleChange = (event: any) => {
    SetMenuname(event.target.value as string);
  };

  const Ingredient_handleChange = (event: any) => {
    SetIngredient(event.target.value as string);
  };

  const calorie_handleChange = (event: any) => {
    SetCalories(event.target.value as number);
  };

  const AddedtimehandleChange = (event: any) => {
    SetAddedTime(event.target.value as string);
  };


  const [menuname, SetMenuname] = useState(String);
  const [calorie, SetCalories] = useState(Number);
  let calories = Number(calorie);

  const [ingredient, SetIngredient] = useState(String);
  const [addedtimes, SetAddedTime] = useState(String);
  const [userID, SetOwnerID] = useState(Number);
  const [menutypeID, SetTypeID] = useState(Number);
  const [menugroupID, SetGroupID] = useState(Number);

  const CreateMenu = async () => {
    const menu = {
      menuname: menuname,
      typeID: menutypeID,
      groupID: menugroupID,
      calories,
      userID: userID,
      addedtime: addedtimes + ":00+07:00",
      ingredient: ingredient,
    };
    console.log(menu);
    const res: any = await api.createMenu({ menu: menu });

    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  //-------------------------------------------------------------------------------------
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`Create Menu`}></Header>

      <Content>
        <Container component="main" maxWidth="xs">
          <form className={classes.form} noValidate autoComplete="off">
            <FormControl className={classes.formControl}>

              <TextField
                id="menuname"
                label="Menu name"
                size="medium"
                //type="string"
                value={menuname}
                onChange={Menuname_handleChange}
              />
            </FormControl>

            <FormControl className={classes.formControl}>
              <InputLabel>Menu Type</InputLabel>
              <Select
                labelId="menutype-label"
                id="type"
                name="menutype"
                value={menutypeID}
                onChange={MenutypeIDhandleChange}
                style={{ width: 400 }}
              >
                {menutypes.map((item: EntMenutype) => (
                  <MenuItem value={item.id}>{item.type}</MenuItem>
                ))}
              </Select>
            </FormControl>

            <FormControl className={classes.formControl}>
              <InputLabel>Menu Group</InputLabel>
              <Select
                labelId="menugroup-label"
                id="group"
                name="menugroup"
                type="string"
                value={menugroupID}
                onChange={MenugroupIDhandleChange}
                style={{ width: 400 }}
              >
                {menugroups.map((item: EntMenugroup) => (
                  <MenuItem value={item.id}>{item.group}</MenuItem>
                ))}
              </Select>
            </FormControl>

            <FormControl className={classes.formControl}>
              <TextField
                id="standard-basic"
                label="Calories"
                size="medium"
                //type="number"
                value={calories}
                onChange={calorie_handleChange}
              />
            </FormControl>

            <FormControl className={classes.formControl}>
              <TextField
                id="standard-basic"
                label="Nutrition and Ingredient"
                size="medium"
                value={ingredient}
                onChange={Ingredient_handleChange}
              />
            </FormControl>

            <FormControl className={classes.formControl}>
              <InputLabel>User Email</InputLabel>
              <Select
                labelId="user-label"
                id="user"
                value={userID}
                onChange={UserIDhandleChange}
                style={{ width: 400 }}
              >
                {users.map((item: EntUser) => (
                  <MenuItem value={item.id}>{item.email}</MenuItem>
                ))}
              </Select>
            </FormControl>

            <FormControl className={classes.formControl}>
              <TextField
                id="addedtime"
                label="Month/Day/Year Added Time"
                type="datetime-local"
                value={addedtimes}
                onChange={AddedtimehandleChange}
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
              />
            </FormControl>

            <div className={classes.root}>
              <Button variant="contained" color="primary" size="large"
                className={classes.button}
                onClick={() => {
                  CreateMenu();
                }}
              >
                Save Menu
              </Button>
              <Link component={RouterLink} to="/">
                <Button variant="contained" color="primary" size="large"
                  className={classes.button}>
                  Back
                </Button>
              </Link>
            </div>
            {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    This is a success alert — check it out!
                  </Alert>
                ) : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      This is a warning alert — check it out!
                    </Alert>
                  )}
              </div>
            ) : null}
          </form>
        </Container>
      </Content>
    </Page>
  );
}