import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntMenu } from '../../api/models/EntMenu';
import  moment  from 'moment'; 

 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [menus, setMenu] = useState<EntMenu[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getMenu = async () => {
     const res = await api.listMenu({ limit: 10, offset: 0 });
     setLoading(false);
     setMenu(res);
   };
   getMenu();
 }, [loading]);
 
 const deleteMenu = async (id: number) => {
   const res = await api.deleteMenu({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">Menu Name</TableCell>
           <TableCell align="center">Type</TableCell>
           <TableCell align="center">Group</TableCell>
           <TableCell align="center">Calories(Kcal)</TableCell>
           <TableCell align="center">Ingredient</TableCell>
           <TableCell align="center">User</TableCell>
           <TableCell align="center">Added Time</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {menus.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.menuname}</TableCell>
             <TableCell align="center">{item.edges?.type?.type}</TableCell>
             <TableCell align="center">{item.edges?.group?.group}</TableCell>
             <TableCell align="center">{item.calories}</TableCell>
             <TableCell align="center">{item.ingredient}</TableCell>
             <TableCell align="center">{item.edges?.owner?.name}</TableCell>
             <TableCell align="center">{moment(item.addedTime).format('MM/DD/YYYY HH.mm')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteMenu(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
