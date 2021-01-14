import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ระบบบันทึกเมนูอาหาร' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`${profile.givenName || 'to Backstage'}`}
     ></Header>
     <Content>
       <ContentHeader title="Application CRUD">
         <Link component={RouterLink} to="/foodmenu">
           <Button variant="contained" color="primary">
             Add Menu
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;

