import React from "react";
import { View, Text, StyleSheet, Dimensions } from "react-native";
import { Header } from "@/components/header";
import {BarChart} from 'react-native-gifted-charts';

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

  const renderTitle = () => {
    return(
      <View
        style={{
          marginBottom: 20,
          flex: 1,
          flexDirection: 'row',
          justifyContent: 'space-between',
        }}>
        <View style={{flexDirection: 'row', alignItems: 'center'}}>
          <Text
            style={{
              fontSize: 18,
              height: 22,
              fontWeight: 'bold',
              color: '#000000',
            }}>
            Daily in Room A
          </Text>
        </View>
        <View style={{
          flexDirection: 'row',
          
        }}>
          <View style={{flexDirection: 'row', alignItems: 'center'}}>
            <View
              style={{
                height: 12,
                width: 12,
                borderRadius: 6,
                backgroundColor: '#34A853',
                marginRight: 8,
              }}
            />
            <Text
              style={{
                width: 30,
                height: 16,
                color: '#000000',
              }}>
              In
            </Text>
          </View>
          <View style={{flexDirection: 'row', alignItems: 'center'}}>
            <View
              style={{
                height: 12,
                width: 12,
                borderRadius: 6,
                backgroundColor: '#ED4A4A',
                marginRight: 8,
              }}
            />
            <Text
              style={{
                width: 30,
                height: 16,
                color: '#000000',
              }}>
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
        <View
          style={{
            backgroundColor: 'white',
            padding: 20,
            paddingTop: 30,
            borderRadius: 10,
            shadowOffset: {
              width: 0,
              height: 0,
            },
            shadowOpacity: 0.2,
            shadowRadius: 5,
          }}>
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
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#FFFFFF",
  },
  header_container: {

  },
  content_container: {
    paddingTop: 30,
    paddingHorizontal: 0.05 * screenWidth,
  }
});
