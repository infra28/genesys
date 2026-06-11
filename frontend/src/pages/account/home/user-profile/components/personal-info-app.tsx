
import { SquarePen } from 'lucide-react';
import { Link } from 'react-router';

import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Table, TableBody, TableCell, TableRow } from '@/components/ui/table';


const PersonalInfoApp = () => {
  return (
    <Card className="min-w-full">
      <CardHeader>
        <CardTitle>Dapodik</CardTitle>
      </CardHeader>
      <CardContent className="kt-scrollable-x-auto pb-3 p-0">
        <Table className="align-middle text-sm text-muted-foreground">
          <TableBody>
            <TableRow>
              <TableCell className="py-2 text-secondary-foreground font-normal">
                Nama Lengkap
              </TableCell>
              <TableCell className="py-2 text-foreground font-normaltext-sm">
                Jason Tatum
              </TableCell>
              <TableCell className="py-2 text-center">
                <Button variant="ghost" mode="icon">
                  <SquarePen size={16} className="text-blue-500" />
                </Button>
              </TableCell>
            </TableRow>

            <TableRow>
              <TableCell className="py-3 text-secondary-foreground font-normal">
                Birthday
              </TableCell>
              <TableCell className="py-3 text-secondary-foreground text-sm font-normal">
                28 May 1996
              </TableCell>
              <TableCell className="py-3 text-center">
                <Button variant="ghost" mode="icon">
                  <SquarePen size={16} className="text-blue-500" />
                </Button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="py-3 text-secondary-foreground font-normal">
                Gender
              </TableCell>
              <TableCell className="py-3 text-secondary-foreground text-sm font-normal">
                Male
              </TableCell>
              <TableCell className="py-3 text-center">
                <Button variant="ghost" mode="icon">
                  <SquarePen size={16} className="text-blue-500" />
                </Button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="py-3">Address</TableCell>
              <TableCell className="py-3 text-secondary-foreground text-sm font-normal">
                You have no an address yet
              </TableCell>
              <TableCell className="py-3 text-center">
                <Button mode="link" underlined="dashed" asChild>
                  <Link to="#">Add</Link>
                </Button>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  );
};

export { PersonalInfoApp };
