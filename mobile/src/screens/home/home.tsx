import React from "react";
import { View, Text, StyleSheet, Dimensions, FlatList } from "react-native";
import { Header } from "@/components/header";
import {BarChart} from 'react-native-gifted-charts';
import { Colors } from "@/theme/variables";
import { autoBatchEnhancer } from "@reduxjs/toolkit";

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const Home = () => {
  const barData = [
    {
      value: 30,
      label: '20',
      spacing: 2,
      labelWidth: 30,
      labelTextStyle: {color: 'gray'},
      frontColor: '#34A853',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30, 
      frontColor: '#ED4A4A',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30,
      label: '21',
      spacing: 2,
      labelWidth: 30,
      labelTextStyle: {color: 'gray'},
      frontColor: '#34A853',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30, 
      frontColor: '#ED4A4A',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30,
      label: '22',
      spacing: 2,
      labelWidth: 30,
      labelTextStyle: {color: 'gray'},
      frontColor: '#34A853',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30, 
      frontColor: '#ED4A4A',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30,
      label: '23',
      spacing: 2,
      labelWidth: 30,
      labelTextStyle: {color: 'gray'},
      frontColor: '#34A853',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30, 
      frontColor: '#ED4A4A',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30,
      label: '24',
      spacing: 2,
      labelWidth: 30,
      labelTextStyle: {color: 'gray'},
      frontColor: '#34A853',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30, 
      frontColor: '#ED4A4A',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30,
      label: '25',
      spacing: 2,
      labelWidth: 30,
      labelTextStyle: {color: 'gray'},
      frontColor: '#34A853',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30, 
      frontColor: '#ED4A4A',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30,
      label: '26',
      spacing: 2,
      labelWidth: 30,
      labelTextStyle: {color: 'gray'},
      frontColor: '#34A853',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
    {
      value: 30, 
      frontColor: '#ED4A4A',
      topLabelComponent: () => (
        <Text style={{color: 'gray', fontSize: 10, marginBottom: 2}}>30</Text>
      ),
    },
  ];

  const DATA = [
    { id: '1', room: 'room 1', in: 10, out: 100, total: 100 },
    { id: '2', room: 'room 2', in: 10, out: 20, total: 30 },
    { id: '3', room: 'room 3', in: 10, out: 20, total: 30 },
    // Add more items as needed
  ];

  type RoomProps = {room: string, in: number, out: number, total: number};

  const Room = (Room: RoomProps) => (
    <View style={styles.room}>
      <Text style={styles.title}>{Room.room}</Text>
      <View style={{ flexDirection: 'column' ,alignItems: 'center'}}>
        <View style={styles.total_card}>
          <Text style={{fontWeight: 'bold', color: '#306fbb'}}>Total: {Room.total}</Text>
        </View >
        <View style={styles.in_card}>
          <Text style={{fontWeight: 'bold', color: '#47b063'}}>In: {Room.in}</Text>
        </View>
        <View style={styles.out_card}>
          <Text style={{fontWeight: 'bold', color: '#ef5e5e'}}>Out: {Room.out}</Text>
        </View>
      </View>
    </View>
  );

  const renderTitle = () => {
    return(
      <View style={styles.barchart_title}>
        <View style={{flexDirection: 'row', alignItems: 'center'}}>
          <Text style={styles.title}>
            Daily in Room A
          </Text>
        </View>
        <View style={{ flexDirection: 'row' }}>
          <View style={{flexDirection: 'row', alignItems: 'center'}}>
            <View style={styles.in_status}/>
            <Text style={styles.text}>
              In
            </Text>
          </View>
          <View style={{flexDirection: 'row', alignItems: 'center'}}>
            <View style={styles.out_status}/>
            <Text style={styles.text}>
              Out
            </Text>
          </View>
        </View>
      </View>
    )
  }

  return (
    <View style={styles.container}>
      <View style={styles.header_container}>
        <Header title="Home" />
      </View>
      <View style={styles.content_container}>
        <View style={styles.component_content}>
          {renderTitle()}
          <BarChart
            data={barData}
            height={150}
            barWidth={14}
            spacing={14}
            roundedTop
            roundedBottom
            hideRules
            xAxisThickness={0}
            yAxisThickness={-7}
            yAxisTextStyle={{color: 'gray'}}
            noOfSections={5}
            maxValue={75}
          />
        </View>
        <View style={styles.component_content}>
          <FlatList
            data={DATA}
            horizontal={true}
            renderItem={({item}) => <Room room={item.room} in={item.in} out={item.out} total={item.total} />}
            keyExtractor={item => item.id}
          />
        </View>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f4f7fb',
  },
  header_container: {
    shadowOffset: {
      width: 0,
      height: 0,
    },
    shadowOpacity: 0.2,
    shadowRadius: 5,
  },
  barchart_title: {
    marginBottom: 20,
    flex: 1,
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  in_status: {
    height: 12,
    width: 12,
    borderRadius: 6,
    backgroundColor: '#34A853',
    marginRight: 8,
  },
  out_status: {
    height: 12,
    width: 12,
    borderRadius: 6,
    backgroundColor: '#ED4A4A',
    marginRight: 8,
  },
  text: {
    width: 30,
    height: 16,
    color: '#000000',
  },
  content_container: {
    paddingTop: 30,
    paddingHorizontal: 0.05 * screenWidth,
  },
  component_content: {
    backgroundColor: 'white',
    padding: 20,
    paddingTop: 30,
    borderRadius: 10,
    marginBottom: 20,
    shadowOffset: {
      width: 0,
      height: 0,
    },
    shadowOpacity: 0.2,
    shadowRadius: 5,
  },
  room: {
    backgroundColor: Colors.WHITE,
    alignItems: 'center',
    padding: 10,
    width: 100,
    marginRight: 5,
    borderRadius: 10,
    borderWidth: 1,
    borderColor: '#7c7c7c',
  },
  total_card: {
    backgroundColor: '#d9e4f3', 
    paddingTop: 10, 
    paddingBottom: 10, 
    alignItems: 'center',
    width: 80,
    borderTopStartRadius: 10,
    borderTopRightRadius: 10,
  },
  in_card: {
    width: 80,
    backgroundColor: '#e0f2e5', 
    paddingTop: 10, 
    paddingBottom: 10, 
    alignItems: 'center', 
  },
  out_card: {
    backgroundColor: '#fde4e4', 
    paddingTop: 10, 
    paddingBottom: 10, 
    alignItems: 'center', 
    width: 80,
    borderBottomLeftRadius: 10,
    borderBottomRightRadius: 10,
  },
  title: {
    fontSize: 18,
    height: 22,
    fontWeight: 'bold',
    color: '#000000',
  },
});
